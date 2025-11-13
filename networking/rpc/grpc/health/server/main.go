package main

import (
	"context"
	"fmt"
	"log"
	"net"

	foov1 "github.com/mukappalambda/go-examples/grpc-health/gen/foo/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
)

type fooServer struct {
	foov1.UnsafeFooServiceServer
}

func (f *fooServer) UnaryFoo(_ context.Context, req *foov1.UnaryFooRequest) (*foov1.UnaryFooResponse, error) {
	return &foov1.UnaryFooResponse{Message: fmt.Sprintf("Received: %s", req.GetMessage())}, nil
}

var _ foov1.FooServiceServer = (*fooServer)(nil)

func main() {
	server := grpc.NewServer()
	healthcheck := health.NewServer()
	healthgrpc.RegisterHealthServer(server, healthcheck)
	foov1.RegisterFooServiceServer(server, &fooServer{})

	var lc net.ListenConfig
	ln, err := lc.Listen(context.Background(), "tcp", "127.0.0.1:9090")
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	if err := server.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
