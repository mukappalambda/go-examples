package services

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/mukappalambda/go-examples/networking/rpc/grpc/note_server/api/services/notes"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

type service struct {
	pb.UnimplementedNoteServiceServer
	notes []*pb.Note
}

func NewServer() *service {
	return &service{notes: make([]*pb.Note, 0)}
}

func (s *service) GetNote(_ context.Context, req *pb.GetNoteRequest) (*pb.GetNoteResponse, error) {
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

func (s *service) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (*pb.CreateNoteResponse, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	note := req.GetNote()
	s.notes = append(s.notes, note)
	log.Printf("[client@%s] created a note (id: %q title: %q) at %s\n", p.Addr.String(), note.GetId(), note.GetTitle(), note.CreatedAt.AsTime())
	return &pb.CreateNoteResponse{Note: note}, nil
}

func (s *service) List(_ context.Context, _ *pb.ListNotesRequest) (*pb.ListNotesResponse, error) {
	return &pb.ListNotesResponse{Notes: s.notes}, nil
}

func StupidInterceptor(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
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
