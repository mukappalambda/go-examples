package main

import (
	"fmt"
	"net"
	"os"

	pb "github.com/mukappalambda/go-examples/networking/rpc/grpc/note_server/api/services/notes"
	"github.com/mukappalambda/go-examples/networking/rpc/grpc/note_server/cmd/noted/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/orca"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	port := 50051
	addr := fmt.Sprintf("localhost:%d", port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("error listening on %s", addr)
	}
	defer ln.Close()
	serverOpts := []grpc.ServerOption{
		orca.CallMetricsServerOption(nil),
		grpc.UnaryInterceptor(services.StupidInterceptor),
	}
	grpcServer := grpc.NewServer(serverOpts...)
	healthgrpc.RegisterHealthServer(grpcServer, health.NewServer())
	pb.RegisterNoteServiceServer(grpcServer, services.NewServer())
	fmt.Printf("server running on %s\n", addr)
	if err := grpcServer.Serve(ln); err != nil {
		return fmt.Errorf("error serving: %v", err)
	}
	return nil
}
