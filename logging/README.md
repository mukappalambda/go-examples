# Logging Demo

`slog`

```bash
go run slog/main.go
```

```bash
# In another terminal
curl http://localhost:8080/name\?\=alpha\&email\=alex@gmail.com
```

Expected output:

```bash
{"time":"2024-11-10T12:14:26.789519171+08:00","level":"INFO","msg":"My-Server","Method":"GET","Path":"/name","Query-String":{"":["alpha"],"email":["alex@gmail.com"]},"User-Agent":"curl/7.68.0","Proto":"HTTP/1.1"}
```

---
