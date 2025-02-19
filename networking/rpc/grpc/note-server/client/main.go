package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/mukappalambda/go-examples/networking/rpc/grpc/note_server/note"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"
)

var (
	defaultTimeout = time.Second
	cancelit       = flag.Bool("cancelit", false, "cancel the client request")
)

func fooInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	elapsed := time.Since(start)
	log.Printf("request start time: %s; request elapsed time: %s\n", start.Format(time.DateTime), elapsed.String())
	return err
}

func barInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func main() {
	flag.Parse()
	port := 50051
	addr := fmt.Sprintf("localhost:%d", port)
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(fooInterceptor),
		grpc.WithStreamInterceptor(barInterceptor),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                defaultTimeout,
			Timeout:             defaultTimeout,
			PermitWithoutStream: true,
		}),
	}
	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		log.Fatalf("error creating client connection: %+v\n", err)
	}
	defer conn.Close()
	client := pb.NewNoteServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if *cancelit {
		cancel()
	}

	if err := createFakeNotes(ctx, client); err != nil {
		log.Fatal(err)
	}
	if err := listNotes(ctx, client); err != nil {
		log.Fatal(err)
	}
}

func listNotes(ctx context.Context, client pb.NoteServiceClient) error {
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

func createFakeNotes(ctx context.Context, client pb.NoteServiceClient) error {
	fakeNotes := []string{"alpha", "beta", "gamma", "delta"}
	for _, fakeNote := range fakeNotes {
		req := &pb.CreateNoteRequest{
			Title:   fakeNote,
			Content: fakeNote,
		}
		res, err := client.CreateNote(ctx, req)
		if err != nil {
			errorCode := status.Code(err)
			if errorCode == codes.Canceled {
				return fmt.Errorf("request got canceled: %s", errorCode)
			}
			return fmt.Errorf("error creating note: %v", err)
		}
		fmt.Printf("response: %+v\n", res)
	}
	return nil
}
