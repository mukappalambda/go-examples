package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/mukappalambda/go-examples/networking/rpc/grpc/note_server/api/services/notes"

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
	notes []*pb.Note
}

func newServer() *server {
	return &server{notes: make([]*pb.Note, 0)}
}

func (s *server) GetNote(_ context.Context, req *pb.GetNoteRequest) (*pb.GetNoteResponse, error) {
	id := req.GetId()
	var note *pb.Note
	for _, n := range s.notes {
		if n.GetId() == id {
			note = n
			break
		}
	}
	if note == nil {
		return nil, fmt.Errorf("note not found")
	}
	return &pb.GetNoteResponse{Note: note}, nil
}

func (s *server) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (*pb.CreateNoteResponse, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	note := req.GetNote()
	s.notes = append(s.notes, note)
	log.Printf("[client@%s] created a note (id: %q title: %q) at %s\n", p.Addr.String(), note.GetId(), note.GetTitle(), note.CreatedAt.AsTime())
	return &pb.CreateNoteResponse{Note: note}, nil
}

func (s *server) List(_ context.Context, _ *pb.ListNotesRequest) (*pb.ListNotesResponse, error) {
	return &pb.ListNotesResponse{Notes: s.notes}, nil
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
	opts := []grpc.ServerOption{
		orca.CallMetricsServerOption(nil),
		grpc.UnaryInterceptor(stupidInterceptor),
	}
	grpcServer := grpc.NewServer(opts...)
	healthgrpc.RegisterHealthServer(grpcServer, health.NewServer())
	pb.RegisterNoteServiceServer(grpcServer, newServer())
	fmt.Printf("server running on %s\n", addr)
	if err := grpcServer.Serve(ln); err != nil {
		return fmt.Errorf("error serving: %v", err)
	}
	return nil
}
