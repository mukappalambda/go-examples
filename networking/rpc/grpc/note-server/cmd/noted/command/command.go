package command

import (
	"fmt"
	"net"

	"github.com/mukappalambda/go-examples/networking/rpc/grpc/note_server/cmd/noted/server"
)

type app struct{}

func App() *app {
	return &app{}
}

func (a *app) Run() error {
	port := 50051
	addr := fmt.Sprintf("localhost:%d", port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("error listening on %s", addr)
	}
	defer ln.Close()
	server, err := server.New()
	if err != nil {
		return err
	}
	fmt.Printf("server running on %s\n", addr)
	if err := server.Serve(ln); err != nil {
		return fmt.Errorf("error serving: %v", err)
	}
	return nil
}
