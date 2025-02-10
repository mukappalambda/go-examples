package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"

	"github.com/mukappalambda/go-examples/networking/rpc/kvstore/shared"
)

type StoreService struct {
	Data map[string]string
}

func NewStoreService() *StoreService {
	ss := &StoreService{
		Data: make(map[string]string),
	}
	return ss
}

func (ss *StoreService) Get(args shared.Args, reply *shared.Reply) error {
	k := args.Key
	v, ok := ss.Data[k]
	if !ok {
		return fmt.Errorf("key %q does not exist", k)
	}
	reply.Value = v
	return nil
}

func (ss *StoreService) Set(args shared.Args, reply *shared.Reply) error {
	k := args.Key
	v := args.Value
	ss.Data[k] = v
	return nil
}

func (ss *StoreService) Delete(args shared.Args, reply *shared.Reply) error {
	k := args.Key
	_, ok := ss.Data[k]
	if !ok {
		reply.Value = ""
		return fmt.Errorf("key %q does not exist", k)
	}
	delete(ss.Data, k)
	return nil
}

func main() {
	addr := "127.0.0.1:8080"
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error listening on %s\n", err)
	}
	defer ln.Close()
	rpcServer := rpc.NewServer()
	err = rpcServer.Register(NewStoreService())
	if err != nil {
		log.Printf("Error registering service: %s\n", err)
		ln.Close()
		os.Exit(1)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s\n", err)
			continue
		}
		go func(conn net.Conn) {
			defer conn.Close()
			rpcServer.ServeConn(conn)
		}(conn)
	}
}
