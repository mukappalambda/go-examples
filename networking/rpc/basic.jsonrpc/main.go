package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	Name   string
	Age    int
	Active bool
	Score  float64
}

type Reply struct {
	Output string
}

type Foo struct{}

func (f *Foo) Run(args Args, reply *Reply) error {
	reply.Output = fmt.Sprintf("%+v", args)
	return nil
}

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatalf("error listening: %s\n", err)
	}
	defer ln.Close()
	serverAddr := ln.Addr().String()
	server := rpc.NewServer()
	if err := server.Register(new(Foo)); err != nil {
		log.Fatalf("error registering service: %s\n", err)
	}
	go func(ln net.Listener) {
		for {
			conn, err := ln.Accept()
			if err != nil {
				ln.Close()
				return
			}
			go func(conn io.ReadWriteCloser) {
				codec := jsonrpc.NewServerCodec(conn)
				server.ServeCodec(codec)
			}(conn)
		}
	}(ln)
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatalf("error connecting to the network: %s\n", err)
	}
	codec := jsonrpc.NewClientCodec(conn)
	client := rpc.NewClientWithCodec(codec)
	defer client.Close()
	args := Args{
		Name:   "alpha",
		Age:    30,
		Active: true,
		Score:  12.34,
	}
	reply := Reply{}
	if err := client.Call("Foo.Run", args, &reply); err != nil {
		log.Fatalf("error invoking the named function: %s\n", err)
	}
	fmt.Printf("%+v\n", reply)
}
