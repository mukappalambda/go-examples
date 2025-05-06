package client

import (
	"context"
	"fmt"
	"log"
	"time"

	notesapi "github.com/mukappalambda/go-examples/networking/rpc/grpc/note_server/api/services/notes"
	"github.com/mukappalambda/go-examples/networking/rpc/grpc/note_server/core/notes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var DefaultTimeout = time.Second

type ClosableGRPCConn interface {
	grpc.ClientConnInterface
	Close() error
}

type Client interface {
	Conn() ClosableGRPCConn
	Connector() func() (*grpc.ClientConn, error)
	CreateNotes(ctx context.Context, notes ...notes.Note) error
	ListNotes(ctx context.Context) error
	NoteService() notes.Store
}

type client struct {
	services
	conn      *grpc.ClientConn
	connector func() (*grpc.ClientConn, error)
}

func NewClient(address string, opts ...grpc.DialOption) (Client, error) {
	c := &client{}
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

func (c *client) Conn() ClosableGRPCConn {
	return c.conn
}

func (c *client) Connector() func() (*grpc.ClientConn, error) {
	return c.connector
}

func (c *client) CreateNotes(ctx context.Context, notes ...notes.Note) error {
	for _, note := range notes {
		n, err := c.NoteService().Create(ctx, note)
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

func (c *client) ListNotes(ctx context.Context) error {
	notes, err := c.NoteService().List(ctx)
	if err != nil {
		return fmt.Errorf("failed to list notes: %s", err)
	}
	for _, note := range notes {
		fmt.Printf("%+v\n", note)
	}
	return nil
}

func (c *client) NoteService() notes.Store {
	if c.noteStore != nil {
		return c.noteStore
	}
	return NewRemoteNoteStore(notesapi.NewNoteServiceClient(c.conn))
}

var _ Client = (*client)(nil)

func FooInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	elapsed := time.Since(start)
	log.Printf("request start time: %s; request elapsed time: %s\n", start.Format(time.DateTime), elapsed.String())
	return err
}

func BarInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	return s, nil
}
