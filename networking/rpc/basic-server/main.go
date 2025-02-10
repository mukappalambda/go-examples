package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"time"
)

type Feat struct{}

func (f *Feat) Compute(args Args, reply *Reply) error {
	startTime, err := time.Parse(time.DateTime, args.StartTime)
	if err != nil {
		return fmt.Errorf("error parsing StartTime: %q", args.StartTime)
	}
	endTime, err := time.Parse(time.DateTime, args.EndTime)
	if err != nil {
		return fmt.Errorf("error parsing EndTime: %q", args.EndTime)
	}
	duration := endTime.Sub(startTime)
	reply.Duration = duration.String()
	return nil
}

type Args struct {
	StartTime string
	EndTime   string
}

type Reply struct {
	Duration string
}

func main() {
	port := 8080
	addr := fmt.Sprintf(":%d", port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error listening on %s\n", addr)
	}
	defer ln.Close()

	newServer := rpc.NewServer()
	err = newServer.Register(new(Feat))
	if err != nil {
		log.Fatalf("Error registering methods: %s\n", err)
	}

	go func() {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			ln.Close()
			os.Exit(1)
		}
		go func(conn net.Conn) {
			defer conn.Close()
			newServer.ServeConn(conn)
		}(conn)
	}()

	serverAddr := ln.Addr().String()
	log.Println("Server listening on", serverAddr)

	startTime := time.Now()
	duration := 10 * time.Minute
	endTime := startTime.Add(duration)

	args := &Args{
		StartTime: startTime.Format(time.DateTime),
		EndTime:   endTime.Format(time.DateTime),
	}
	reply := new(Reply)

	client, err := rpc.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatalf("Error establishing connection: %s\n", err)
	}
	defer client.Close()

	err = client.Call("Feat.Compute", args, reply)
	if err != nil {
		log.Println(err)
		client.Close()
		os.Exit(1)
	}
	fmt.Printf("Reply: %+v\n", reply)

	args = &Args{
		StartTime: "incorrect start time format",
		EndTime:   endTime.Format(time.DateTime),
	}
	reply = new(Reply)
	err = client.Call("Feat.Compute", args, reply)
	if err != nil {
		log.Println(err)
	}

	args = &Args{
		StartTime: startTime.Format(time.DateTime),
		EndTime:   "incorrect end time format",
	}
	reply = new(Reply)
	err = client.Call("Feat.Compute", args, reply)
	if err != nil {
		log.Println(err)
	}
}
