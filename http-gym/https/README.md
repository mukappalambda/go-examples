# Set up a web server with HTTPS

Generate the server key and server certificate:

```bash
bash run.sh
```

Run the application:

```bash
go run main.go
```

Make a client request:

```bash
curl --insecure https://localhost:4443
```
