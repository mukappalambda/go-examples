# Note Server

Install Go plugins for the protocol compiler if you haven't:

(ref: [gRPC Quick Start](https://grpc.io/docs/languages/go/quickstart/))

Build:

```bash
make build
```

Start the server:

```bash
./bin/noted
```

Start the client:

```bash
./bin/note-cli
```

Generate proto files:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

make protoc
```

Clean binaries:

```bash
make clean
```
