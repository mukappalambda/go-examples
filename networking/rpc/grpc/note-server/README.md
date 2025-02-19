# Note Server

Install Go plugins for the protocol compiler if you haven't:

(ref: [gRPC Quick Start](https://grpc.io/docs/languages/go/quickstart/))

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

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
