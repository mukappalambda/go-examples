package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/mukappalambda/go-examples/networking/rpc/grpc/note_server/api/services/notes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
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
		return err
	}
	if err := listNotes(ctx, client); err != nil {
		return err
	}
	return nil
}

func listNotes(ctx context.Context, client pb.NoteServiceClient) error {
	resp, err := client.List(ctx, &pb.ListNotesRequest{Filters: []string{}})
	if err != nil {
		return fmt.Errorf("failed to list notes: %s", err)
	}
	for _, n := range resp.GetNotes() {
		fmt.Printf("%+v\n", n)
	}
	return nil
}

func createFakeNotes(ctx context.Context, client pb.NoteServiceClient) error {
	fakeNotes := []*pb.Note{
		{
			Id:        "1",
			Title:     "alpha's title",
			Content:   "alpha's content",
			CreatedAt: &timestamppb.Timestamp{Seconds: time.Now().Add(-2 * time.Second).Unix()},
			UpdatedAt: &timestamppb.Timestamp{Seconds: time.Now().Add(-2 * time.Second).Unix()},
		},
		{
			Id:        "2",
			Title:     "beta's title",
			Content:   "beta's content",
			CreatedAt: &timestamppb.Timestamp{Seconds: time.Now().Add(-1 * time.Second).Unix()},
			UpdatedAt: &timestamppb.Timestamp{Seconds: time.Now().Add(-1 * time.Second).Unix()},
		},
		{
			Id:        "3",
			Title:     "gamma's title",
			Content:   "gamma's content",
			CreatedAt: &timestamppb.Timestamp{Seconds: time.Now().Unix()},
			UpdatedAt: &timestamppb.Timestamp{Seconds: time.Now().Unix()},
		},
	}

	for _, note := range fakeNotes {
		req := &pb.CreateNoteRequest{
			Note: note,
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
