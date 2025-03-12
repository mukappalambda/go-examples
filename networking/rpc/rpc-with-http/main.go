package main

import (
	"fmt"
	"net"
	"net/http/httptest"
	"net/rpc"
	"os"
)

type Args struct {
	FirstName string
	LastName  string
}

type Reply struct {
	Result string
}

type Foo struct{}

func (f *Foo) Run(args Args, reply *Reply) error { //nolint
	reply.Result = fmt.Sprintf("%s %s", args.FirstName, args.LastName)
	return nil
}
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	addr := "127.0.0.1:8080"
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("error listening on %q: %w", addr, err)
	}
	defer ln.Close()
	server := rpc.NewServer()
	if err := server.Register(new(Foo)); err != nil && err != rpc.ErrShutdown {
		return err
	}
	go server.Accept(ln)
	server.HandleHTTP("/foo", "/bar")
	httpServer := httptest.NewServer(nil)
	defer httpServer.Close()
	httpServerAddr := httpServer.Listener.Addr().String()
	fmt.Println("httpServerAddr", httpServerAddr)

	client, err := rpc.DialHTTPPath("tcp", httpServerAddr, "/foo")
	if err != nil {
		return err
	}
	defer client.Close()
	args := Args{
		FirstName: "alhpa",
		LastName:  "beta",
	}
	reply := new(Reply)

	if err := client.Call("Foo.Run", args, &reply); err != nil {
		client.Close()
		return err
	}
	fmt.Printf("reply: %+v\n", reply)
	return nil
}
