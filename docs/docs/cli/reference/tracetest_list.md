# CLI Reference
## tracetest list

List resources

### Synopsis

List resources from your Tracetest server

```
tracetest list analyzer|config|datastore|demo|env|organization|pollingprofile|test|testrunner|testsuite|variableset [flags]
```

### Options

```
      --all                    All
  -h, --help                   help for list
      --skip int32             Skip number
      --sortBy string          Sort by
      --sortDirection string   Sort direction (default "desc")
      --take int32             Take number (default 20)
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
