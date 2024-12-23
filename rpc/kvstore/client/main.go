package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/mukappalambda/go-examples/rpc/kvstore/shared"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("Error establishing connection: %s\n", err)
	}
	defer client.Close()
	var args shared.Args
	var reply *shared.Reply

	args = shared.Args{
		Key:   "alpha",
		Value: "1",
	}
	reply = &shared.Reply{}
	err = client.Call("StoreService.Set", args, reply)
	if err != nil {
		log.Println(err)
	}

	args = shared.Args{
		Key:   "beta",
		Value: "2",
	}
	reply = &shared.Reply{}
	err = client.Call("StoreService.Set", args, reply)
	if err != nil {
		log.Println(err)
	}

	args = shared.Args{
		Key:   "beta",
		Value: "3",
	}
	reply = &shared.Reply{}
	err = client.Call("StoreService.Set", args, reply)
	if err != nil {
		log.Println(err)
	}

	args = shared.Args{
		Key: "alpha",
	}
	reply = &shared.Reply{}
	err = client.Call("StoreService.Get", args, reply)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Reply: %+v\n", reply)

	args = shared.Args{
		Key: "beta",
	}
	reply = &shared.Reply{}
	err = client.Call("StoreService.Get", args, reply)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Reply: %+v\n", reply)
}
