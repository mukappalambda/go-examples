package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/mukappalambda/go-examples/networking/rpc/grpc/note_server/note"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/orca"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedNoteServiceServer
}

func (s *server) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (*pb.CreateNoteResponse, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	log.Printf("[client@%s] created a note - title: %q\n", p.Addr.String(), req.GetTitle())
	res := &pb.CreateNoteResponse{
		Id: fmt.Sprintf("id-of-%s", req.GetTitle()),
	}
	return res, nil
}

func (s *server) GetNote(_ context.Context, _ *pb.GetNoteRequest) (*pb.GetNoteResponse, error) {
	return nil, nil
}

func (s *server) ListNotes(_ *pb.ListNotesRequest, _ grpc.ServerStreamingServer[pb.Note]) error {
	return nil
}

func stupidInterceptor(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	p, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("validated [client@%s]", p.Addr.String())
	}
	start := time.Now()
	res, err := handler(ctx, req)
	elapsed := time.Since(start)
	log.Printf("request start time: %s; request elapsed time: %s\n", start.Format(time.DateTime), elapsed.String())
	return res, err
}

func newServer() *server {
	return &server{}
}

func main() {
	port := 50051
	addr := fmt.Sprintf("localhost:%d", port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("error listening on %s\n", addr)
	}
	defer ln.Close()
	opts := []grpc.ServerOption{
		orca.CallMetricsServerOption(nil),
		grpc.UnaryInterceptor(stupidInterceptor),
	}
	grpcServer := grpc.NewServer(opts...)
	healthgrpc.RegisterHealthServer(grpcServer, health.NewServer())
	pb.RegisterNoteServiceServer(grpcServer, newServer())
	if err := grpcServer.Serve(ln); err != nil {
		log.Fatalf("error serving: %v\n", err)
	}
}
