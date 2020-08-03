httpproxy-go-runner
===============

Run https/https proxy separately with [httpproxy-go](https://github.com/a0s/httpproxy-go).

Options
-------

```shell script
  -host string
        bind to host (default "127.0.0.1")
  -http
        enable http proxy
  -https
        enable https proxy
  -port uint
        bind to port (default 8080)
```

Usage
-----

```shell script
go run main.go --host 0.0.0.0 --port 18080 --https
```

or

```shell script
docker run --rm -p 18080:18080 a00s/httpproxy-go-runner --host 0.0.0.0 --port 18080 --https
```
