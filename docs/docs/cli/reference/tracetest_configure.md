# CLI Reference
## tracetest configure

Configure your tracetest CLI

### Synopsis

Configure your tracetest CLI

```
tracetest configure [flags]
```

### Options

```
      --ci                    if cloud is used, don't ask for authentication
      --environment string    set environmentID, so the CLI won't ask you for it
  -g, --global                configuration will be saved in your home dir
  -h, --help                  help for configure
      --organization string   set organizationID, so the CLI won't ask you for it
  -t, --token string          set authetication with token, so the CLI won't ask you for authentication
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
