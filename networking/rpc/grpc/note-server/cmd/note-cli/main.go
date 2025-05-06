package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/mukappalambda/go-examples/networking/rpc/grpc/note_server/client"
	"github.com/mukappalambda/go-examples/networking/rpc/grpc/note_server/core/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
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
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(client.FooInterceptor),
		grpc.WithStreamInterceptor(client.BarInterceptor),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                client.DefaultTimeout,
			Timeout:             client.DefaultTimeout,
			PermitWithoutStream: true,
		}),
	}
	c, err := client.NewClient(addr, opts...)
	if err != nil {
		return fmt.Errorf("error creating client connection: %+v", err)
	}
	conn := c.Conn()
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if *cancelit {
		cancel()
	}

	if err := c.CreateNotes(ctx, store.FakeNotes...); err != nil {
		return err
	}

	if err := c.ListNotes(ctx); err != nil {
		return err
	}
	return nil
}
