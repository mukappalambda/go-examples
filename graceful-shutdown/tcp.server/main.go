package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"os/signal"
	"syscall"
)

var buf = make([]byte, 1024)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("error listening on the network address: %s\n", err)
	}
	go onCtxDone(ctx, ln)

	log.Printf("server listening on %s\n", ln.Addr())

	for {
		conn, err := ln.Accept()
		if err != nil {
			break
		}
		go handleConnection(conn)
	}
}

func onCtxDone(ctx context.Context, ln net.Listener) {
	<-ctx.Done()
	if err := ln.Close(); err != nil {
		log.Printf("error closing the listener: %s\n", err)
	}
	log.Println("listener is shut down gracefully.")
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				continue
			}
			break
		}
		log.Print(string(buf[:n]))
	}
}
