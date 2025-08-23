# HTTP slog handler

Build the binary:

```console
$ go build
```

Print the help message:

```console
$ ./http-slog-handler -h
Usage of ./http-slog-handler:
  -addr string
        server address (default ":8080")
  -readheadertimeout duration
        server readheadertimeout (default 500ms)
  -readtimeout duration
        server readtimeout (default 500ms)
```

Start the server:

```console
$ ./http-slog-handler
Server is listening on :8080
```

Make an HTTP request to the server:

```console
$ curl localhost:8080/
```

Server's logging:

```bash
{"time":"2025-08-23T16:48:33.467773187+08:00","level":"INFO","msg":"My-Server","Method":"GET","Path":"/","Query-String":{},"User-Agent":"curl/8.5.0","Proto":"HTTP/1.1"}
```
