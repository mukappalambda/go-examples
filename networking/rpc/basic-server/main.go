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
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	port := 8080
	addr := fmt.Sprintf("127.0.0.1:%d", port)
	var err error
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("error listening on %q: %w", addr, err)
	}
	defer ln.Close()

	newServer := rpc.NewServer()
	err = newServer.Register(new(Feat))
	if err != nil {
		return fmt.Errorf("error registering methods: %w", err)
	}

	go func() {
		conn, e := ln.Accept()
		err = e
		go func(conn net.Conn) {
			defer conn.Close()
			newServer.ServeConn(conn)
		}(conn)
	}()
	if err != nil {
		ln.Close()
		return fmt.Errorf("failed to accept connection: %w", err)
	}

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
		return fmt.Errorf("error establishing connection: %w", err)
	}
	defer client.Close()

	err = client.Call("Feat.Compute", args, reply)
	if err != nil {
		client.Close()
		return fmt.Errorf("failed to call remote method: %w", err)
	}
	fmt.Printf("Reply: %+v\n", reply)

	args = &Args{
		StartTime: "incorrect start time format",
		EndTime:   endTime.Format(time.DateTime),
	}
	reply = new(Reply)
	err = client.Call("Feat.Compute", args, reply)
	if err != nil {
		client.Close()
		return fmt.Errorf("failed to call remote method: %w", err)
	}

	args = &Args{
		StartTime: startTime.Format(time.DateTime),
		EndTime:   "incorrect end time format",
	}
	reply = new(Reply)
	err = client.Call("Feat.Compute", args, reply)
	if err != nil {
		client.Close()
		return fmt.Errorf("failed to call remote method: %w", err)
	}
	return nil
}
