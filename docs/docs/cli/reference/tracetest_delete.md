# CLI Reference
## tracetest delete

Delete resources

### Synopsis

Delete resources from your Tracetest server

```
tracetest delete analyzer|config|datastore|demo|env|organization|pollingprofile|test|testrunner|testsuite|variableset [flags]
```

### Options

```
  -h, --help        help for delete
      --id string   id of the resource to delete
```

### Options inherited from parent commands

```
  -c, --config string       config file will be used by the CLI (default "config.yml")
  -o, --output string       output format [pretty|json|yaml]
  -s, --server-url string   server url
  -v, --verbose             display debug information
```

### SEE ALSO

* [tracetest](tracetest.md)	 - CLI to configure, install and execute tests on a Tracetest server

###### Auto generated by spf13/cobra on 29-Apr-2024
