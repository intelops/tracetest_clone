/*
TraceTest

OpenAPI definition for TraceTest endpoint and resources

API version: 0.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RunGroupSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunGroupSummary{}

// RunGroupSummary struct for RunGroupSummary
type RunGroupSummary struct {
	Pending *int32 `json:"pending,omitempty"`
	Succeed *int32 `json:"succeed,omitempty"`
	Failed  *int32 `json:"failed,omitempty"`
}

// NewRunGroupSummary instantiates a new RunGroupSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunGroupSummary() *RunGroupSummary {
	this := RunGroupSummary{}
	return &this
}

// NewRunGroupSummaryWithDefaults instantiates a new RunGroupSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunGroupSummaryWithDefaults() *RunGroupSummary {
	this := RunGroupSummary{}
	return &this
}

// GetPending returns the Pending field value if set, zero value otherwise.
func (o *RunGroupSummary) GetPending() int32 {
	if o == nil || isNil(o.Pending) {
		var ret int32
		return ret
	}
	return *o.Pending
}

// GetPendingOk returns a tuple with the Pending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunGroupSummary) GetPendingOk() (*int32, bool) {
	if o == nil || isNil(o.Pending) {
		return nil, false
	}
	return o.Pending, true
}

// HasPending returns a boolean if a field has been set.
func (o *RunGroupSummary) HasPending() bool {
	if o != nil && !isNil(o.Pending) {
		return true
	}

	return false
}

// SetPending gets a reference to the given int32 and assigns it to the Pending field.
func (o *RunGroupSummary) SetPending(v int32) {
	o.Pending = &v
}

// GetSucceed returns the Succeed field value if set, zero value otherwise.
func (o *RunGroupSummary) GetSucceed() int32 {
	if o == nil || isNil(o.Succeed) {
		var ret int32
		return ret
	}
	return *o.Succeed
}

// GetSucceedOk returns a tuple with the Succeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunGroupSummary) GetSucceedOk() (*int32, bool) {
	if o == nil || isNil(o.Succeed) {
		return nil, false
	}
	return o.Succeed, true
}

// HasSucceed returns a boolean if a field has been set.
func (o *RunGroupSummary) HasSucceed() bool {
	if o != nil && !isNil(o.Succeed) {
		return true
	}

	return false
}

// SetSucceed gets a reference to the given int32 and assigns it to the Succeed field.
func (o *RunGroupSummary) SetSucceed(v int32) {
	o.Succeed = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *RunGroupSummary) GetFailed() int32 {
	if o == nil || isNil(o.Failed) {
		var ret int32
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunGroupSummary) GetFailedOk() (*int32, bool) {
	if o == nil || isNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *RunGroupSummary) HasFailed() bool {
	if o != nil && !isNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given int32 and assigns it to the Failed field.
func (o *RunGroupSummary) SetFailed(v int32) {
	o.Failed = &v
}

func (o RunGroupSummary) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunGroupSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Pending) {
		toSerialize["pending"] = o.Pending
	}
	if !isNil(o.Succeed) {
		toSerialize["succeed"] = o.Succeed
	}
	if !isNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	return toSerialize, nil
}

type NullableRunGroupSummary struct {
	value *RunGroupSummary
	isSet bool
}

func (v NullableRunGroupSummary) Get() *RunGroupSummary {
	return v.value
}

func (v *NullableRunGroupSummary) Set(val *RunGroupSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableRunGroupSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableRunGroupSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunGroupSummary(val *RunGroupSummary) *NullableRunGroupSummary {
	return &NullableRunGroupSummary{value: val, isSet: true}
}

func (v NullableRunGroupSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunGroupSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
