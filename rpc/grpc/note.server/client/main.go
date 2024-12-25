package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/mukappalambda/go-examples/rpc/grpc/note.server/note"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var defaultTimeout = time.Second

func main() {
	port := 50051
	addr := fmt.Sprintf("localhost:%d", port)
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		log.Fatalf("error creating client connection: %+v\n", err)
	}
	defer conn.Close()
	client := pb.NewNoteServiceClient(conn)
	req := &pb.CreateNoteRequest{
		Title:   "my-title",
		Content: "my-content",
	}
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	res, err := client.CreateNote(ctx, req)
	if err != nil {
		log.Fatalf("error creating note: %v\n", err)
	}
	fmt.Printf("response: %+v\n", res)
}
