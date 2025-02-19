package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/google/uuid"
	pb "github.com/mukappalambda/go-examples/networking/rpc/grpc/note_server/note"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/orca"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	pb.UnimplementedNoteServiceServer
	notes []*pb.Note
}

func (s *server) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (*pb.CreateNoteResponse, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	noteID := uuid.New().String()
	noteTitle := req.GetTitle()
	noteContent := req.GetContent()
	note := &pb.Note{
		Id:        noteID,
		Title:     noteTitle,
		Content:   noteContent,
		CreatedAt: timestamppb.Now(),
	}
	s.notes = append(s.notes, note)
	log.Printf("[client@%s] created a note (id: %q title: %q) at %s\n", p.Addr.String(), note.GetId(), note.GetTitle(), note.CreatedAt.AsTime())
	return &pb.CreateNoteResponse{Id: noteID}, nil
}

func (s *server) GetNote(_ context.Context, _ *pb.GetNoteRequest) (*pb.GetNoteResponse, error) {
	return nil, nil
}

func (s *server) ListNotes(_ *pb.ListNotesRequest, stream grpc.ServerStreamingServer[pb.Note]) error {
	for _, note := range s.notes {
		if err := stream.Send(note); err != nil {
			return err
		}
	}
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
	return &server{notes: make([]*pb.Note, 0)}
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
	fmt.Printf("server running on %s\n", addr)
	if err := grpcServer.Serve(ln); err != nil {
		log.Fatalf("error serving: %v\n", err)
	}
}
