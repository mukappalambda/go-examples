package main

import (
	"fmt"
	"io"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
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
	if args.Name == "" {
		return fmt.Errorf("invalid name: %q", args.Name)
	}
	if args.Age < 0 {
		return fmt.Errorf("invalid age: %q", args.Age)
	}
	reply.Output = fmt.Sprintf("%+v", args)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return fmt.Errorf("error listening: %w", err)
	}
	defer ln.Close()
	serverAddr := ln.Addr().String()
	server := rpc.NewServer()
	if err := server.Register(new(Foo)); err != nil {
		return fmt.Errorf("error registering service: %w", err)
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
		return fmt.Errorf("error connecting to the network: %w", err)
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
		return fmt.Errorf("error invoking the named function: %w", err)
	}
	fmt.Printf("%+v\n", reply)
	return nil
}
