package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/mukappalambda/go-examples/networking/rpc/grpc/note_server/client"
	"github.com/mukappalambda/go-examples/networking/rpc/grpc/note_server/core/notes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var cancelit = flag.Bool("cancelit", false, "cancel the client request")

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
	client, err := client.New(addr)
	if err != nil {
		return fmt.Errorf("error creating client connection: %+v", err)
	}
	conn := client.Conn()
	defer conn.Close()
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

func listNotes(ctx context.Context, client *client.Client) error {
	notes, err := client.NoteService().List(ctx)
	if err != nil {
		return fmt.Errorf("failed to list notes: %s", err)
	}
	for _, note := range notes {
		fmt.Printf("%+v\n", note)
	}
	return nil
}

func createFakeNotes(ctx context.Context, client *client.Client) error {
	fakeNotes := []notes.Note{
		{
			ID:        "1",
			Title:     "alpha's title",
			Content:   "alpha's content",
			CreatedAt: time.Now().Add(-2 * time.Second),
			UpdatedAt: time.Now().Add(-2 * time.Second),
		},
		{
			ID:        "2",
			Title:     "beta's title",
			Content:   "beta's content",
			CreatedAt: time.Now().Add(-1 * time.Second),
			UpdatedAt: time.Now().Add(-1 * time.Second),
		},
		{
			ID:        "3",
			Title:     "gamma's title",
			Content:   "gamma's content",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, note := range fakeNotes {
		n, err := client.NoteService().Create(ctx, note)
		if err != nil {
			errorCode := status.Code(err)
			if errorCode == codes.Canceled {
				return fmt.Errorf("request got canceled: %s", errorCode)
			}
			return fmt.Errorf("error creating note: %v", err)
		}
		fmt.Printf("response: %+v\n", n)
	}
	return nil
}
