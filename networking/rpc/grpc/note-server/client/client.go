package client

import (
	"context"
	"fmt"
	"log"
	"time"

	notesapi "github.com/mukappalambda/go-examples/networking/rpc/grpc/note_server/api/services/notes"
	"github.com/mukappalambda/go-examples/networking/rpc/grpc/note_server/core/notes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

var defaultTimeout = time.Second

type Client struct {
	services
	conn      *grpc.ClientConn
	connector func() (*grpc.ClientConn, error)
}

func New(address string) (*Client, error) {
	c := &Client{}
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
	connector := func() (*grpc.ClientConn, error) {
		conn, err := grpc.NewClient(address, opts...)
		if err != nil {
			return nil, fmt.Errorf("error creating client connection: %+v", err)
		}
		return conn, nil
	}
	conn, err := connector()
	if err != nil {
		return nil, err
	}
	c.conn, c.connector = conn, connector
	return c, nil
}

func (c *Client) Conn() *grpc.ClientConn {
	return c.conn
}

func (c *Client) NoteService() notes.Store {
	if c.noteStore != nil {
		return c.noteStore
	}
	return NewRemoteNoteStore(notesapi.NewNoteServiceClient(c.conn))
}

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
