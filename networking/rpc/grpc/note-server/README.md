# Note Server

```bash
git clone --depth 1 https://github.com/googleapis/googleapis.git
git clone --depth 1 https://github.com/protocolbuffers/protobuf.git

make build
```

Server:

```bash
go run server/main.go
```

Client:

```bash
go run client/main.go
```
