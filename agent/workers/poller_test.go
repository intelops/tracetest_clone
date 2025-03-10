package workers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/kubeshop/tracetest/agent/client"
	"github.com/kubeshop/tracetest/agent/client/mocks"
	"github.com/kubeshop/tracetest/agent/collector"
	"github.com/kubeshop/tracetest/agent/proto"
	"github.com/kubeshop/tracetest/agent/workers"
	"github.com/kubeshop/tracetest/agent/workers/poller"
	"github.com/kubeshop/tracetest/server/pkg/id"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v1 "go.opentelemetry.io/proto/otlp/trace/v1"
)

func TestPollerWorker(t *testing.T) {
	ctx := ContextWithTracingEnabled()
	controlPlane := mocks.NewGrpcServer()

	client, err := client.Connect(ctx, controlPlane.Addr())
	require.NoError(t, err)

	pollerWorker := workers.NewPollerWorker(client, workers.WithPollerStoppableProcessRunner(workers.NewProcessStopper().RunStoppableProcess))

	client.OnPollingRequest(func(ctx context.Context, pr *proto.PollingRequest) error {
		return pollerWorker.Poll(ctx, pr)
	})

	err = client.Start(ctx)
	require.NoError(t, err)

	tempoAPI := createTempoFakeApi()

	req := proto.PollingRequest{
		Metadata: map[string]string{
			"traceparent": "00-e42e7689e67c64b65ddd3a023a2f8f9d-afad25d95241afce-01",
		},
		TestID:  "test",
		RunID:   1,
		TraceID: "42a2c381da1a5b3a32bc4988bf2431b0",
		Datastore: &proto.DataStore{
			Type: "tempo",
			Tempo: &proto.TempoConfig{
				Type: "http",
				Http: &proto.HttpClientSettings{
					Url: tempoAPI.URL,
				},
			},
		},
	}

	controlPlane.SendPollingRequest(ctx, &req)

	time.Sleep(1 * time.Second)

	// expect traces to be sent to endpoint
	resp := controlPlane.GetLastPollingResponse()
	require.NotNil(t, resp, "agent did not send polling response back to server")

	// Very rudimentar sorting algorithm for only two items in the array
	// first item is always the root span, second is it's child
	var spans = make([]*proto.Span, 2)
	for _, span := range resp.Data.Spans {
		if span.ParentId == "" {
			spans[0] = span
		} else {
			spans[1] = span
		}
	}

	assert.Len(t, spans, 2)
	assert.Equal(t, "", spans[0].ParentId)
	assert.Equal(t, spans[0].Id, spans[1].ParentId)

	assert.Equal(t, req.Metadata["traceparent"], resp.Data.Metadata["traceparent"])
}

func createTempoFakeApi() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
			"batches": [{
				"scopeSpans": [
					{
						"spans": [
							{
								"spanId": "42a2c381da1a5b3a32bc4988bf2431b0",
								"parentSpanId": "",
								"name": "root",
								"kind": "internal",
								"startTimeUnixNano": "0",
								"endTimeUnixNano": "100",
								"attributes": [],
								"events": [],
								"status": {"code": "ok"}
							},
							{
								"spanId": "99a2c381da1a5b3a32bc4988bf2431c3",
								"parentSpanId": "42a2c381da1a5b3a32bc4988bf2431b0",
								"name": "span 2",
								"kind": "internal",
								"startTimeUnixNano": "0",
								"endTimeUnixNano": "100",
								"attributes": [],
								"events": [],
								"status": {"code": "ok"}
							}
						]
					}
				]
			}]
		}`))
		w.WriteHeader(http.StatusOK)
	}))
}

func TestPollerWorkerWithInmemoryDatastore(t *testing.T) {
	ctx := context.Background()
	controlPlane := mocks.NewGrpcServer()

	client, err := client.Connect(ctx, controlPlane.Addr())
	require.NoError(t, err)

	cache := collector.NewTraceCache()

	pollerWorker := workers.NewPollerWorker(client, workers.WithInMemoryDatastore(
		poller.NewInMemoryDatastore(cache),
	),
		workers.WithPollerStoppableProcessRunner(workers.NewProcessStopper().RunStoppableProcess),
	)

	client.OnPollingRequest(func(ctx context.Context, pr *proto.PollingRequest) error {
		return pollerWorker.Poll(ctx, pr)
	})

	err = client.Start(ctx)
	require.NoError(t, err)

	traceID := id.NewRandGenerator().TraceID()
	pollingRequest := proto.PollingRequest{
		TestID:  "test",
		RunID:   1,
		TraceID: traceID.String(),
		Datastore: &proto.DataStore{
			Type: "datadog",
		},
	}

	controlPlane.SendPollingRequest(ctx, &pollingRequest)

	time.Sleep(1 * time.Second)

	// expect traces to not be sent to endpoint
	pollingResponse := controlPlane.GetLastPollingResponse()
	require.NotNil(t, pollingResponse, "agent did not send polling response back to server")

	assert.False(t, pollingResponse.Data.TraceFound)
	assert.Len(t, pollingResponse.Data.Spans, 0)

	span1ID := id.NewRandGenerator().SpanID()
	span2ID := id.NewRandGenerator().SpanID()

	cache.Append(traceID.String(), []*v1.Span{
		{Name: "span 1", ParentSpanId: nil, SpanId: span1ID[:], TraceId: traceID[:]},
		{Name: "span 2", ParentSpanId: span1ID[:], SpanId: span2ID[:], TraceId: traceID[:]},
	})

	controlPlane.SendPollingRequest(ctx, &pollingRequest)

	time.Sleep(1 * time.Second)

	// expect traces to be sent to endpoint
	pollingResponse = controlPlane.GetLastPollingResponse()
	require.NotNil(t, pollingResponse, "agent did not send polling response back to server")

	assert.True(t, pollingResponse.Data.TraceFound)
	assert.Len(t, pollingResponse.Data.Spans, 2)
}

func TestPollerWithInvalidDataStore(t *testing.T) {
	ctx := context.Background()
	controlPlane := mocks.NewGrpcServer()

	client, err := client.Connect(ctx, controlPlane.Addr())
	require.NoError(t, err)

	pollerWorker := workers.NewPollerWorker(client, workers.WithPollerStoppableProcessRunner(workers.NewProcessStopper().RunStoppableProcess))

	client.OnPollingRequest(func(ctx context.Context, pr *proto.PollingRequest) error {
		return pollerWorker.Poll(ctx, pr)
	})

	err = client.Start(ctx)
	require.NoError(t, err)

	pollingRequest := proto.PollingRequest{
		TestID:  "test",
		RunID:   1,
		TraceID: "42a2c381da1a5b3a32bc4988bf2431b0",
		Datastore: &proto.DataStore{
			Type: "tempo",
			Tempo: &proto.TempoConfig{
				Type: "http",
				Http: &proto.HttpClientSettings{
					Url: "http://localhost:12312", // invalid tempo port, this should cause an error
				},
			},
		},
	}

	controlPlane.SendPollingRequest(ctx, &pollingRequest)

	time.Sleep(1 * time.Second)

	pollingResponse := controlPlane.GetLastPollingResponse()
	require.NotNil(t, pollingResponse, "agent did not send polling response back to server")
	require.NotNil(t, pollingResponse.Data.Error)
	assert.Contains(t, pollingResponse.Data.Error.Message, "connection refused")
}
