package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/mukappalambda/go-examples/networking/rpc/grpc/note_server/note"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var defaultTimeout = time.Second

func fooInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	elapsed := time.Since(start)
	log.Printf("request start time: %s; request elapsed time: %s\n", start.Format(time.DateTime), elapsed.String())
	return err
}

func main() {
	port := 50051
	addr := fmt.Sprintf("localhost:%d", port)
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(fooInterceptor),
	}
	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		log.Fatalf("error creating client connection: %+v\n", err)
	}
	defer conn.Close()
	client := pb.NewNoteServiceClient(conn)
	if err := createFakeNotes(client); err != nil {
		log.Fatal(err)
	}
	if err := listNotes(client); err != nil {
		log.Fatal(err)
	}
}

func listNotes(client pb.NoteServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	stream, err := client.ListNotes(ctx, &pb.ListNotesRequest{Page: 0, PageSize: 0})
	if err != nil {
		return fmt.Errorf("failed to list notes: %s", err)
	}
	for {
		note, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("received note: %+v\n", note)
	}
	return nil
}

func createFakeNotes(client pb.NoteServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	fakeNotes := []string{"alpha", "beta", "gamma", "delta"}
	for _, fakeNote := range fakeNotes {
		req := &pb.CreateNoteRequest{
			Title:   fakeNote,
			Content: fakeNote,
		}
		res, err := client.CreateNote(ctx, req)
		if err != nil {
			return fmt.Errorf("error creating note: %v", err)
		}
		fmt.Printf("response: %+v\n", res)
	}
	return nil
}
