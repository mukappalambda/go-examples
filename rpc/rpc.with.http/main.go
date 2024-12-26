package main

import (
	"fmt"
	"log"
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

func (f *Foo) Run(args Args, reply *Reply) error {
	reply.Result = fmt.Sprintf("%s %s", args.FirstName, args.LastName)
	return nil
}

func main() {
	addr := ":8080"
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("error listening on %s\n", addr)
	}
	defer ln.Close()
	server := rpc.NewServer()
	if err := server.Register(new(Foo)); err != nil && err != rpc.ErrShutdown {
		log.Fatal(err)
	}
	go server.Accept(ln)
	server.HandleHTTP("/foo", "/bar")
	httpServer := httptest.NewServer(nil)
	defer httpServer.Close()
	httpServerAddr := httpServer.Listener.Addr().String()
	fmt.Println("httpServerAddr", httpServerAddr)

	client, err := rpc.DialHTTPPath("tcp", httpServerAddr, "/foo")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	args := Args{
		FirstName: "alhpa",
		LastName:  "beta",
	}
	reply := new(Reply)

	if err := client.Call("Foo.Run", args, &reply); err != nil {
		log.Println(err)
		client.Close()
		os.Exit(1)
	}
	fmt.Printf("reply: %+v\n", reply)
}
