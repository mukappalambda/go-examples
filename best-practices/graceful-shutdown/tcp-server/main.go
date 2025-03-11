package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var buf = make([]byte, 1024)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	var err error
	var ln net.Listener
	ln, err = net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		return fmt.Errorf("error listening on the network address: %w", err)
	}
	go func() {
		err = onCtxDone(ctx, ln)
	}()
	if err != nil {
		return fmt.Errorf("error closing the listener: %w", err)
	}

	fmt.Printf("Server is listening on %q\n", ln.Addr())

	for {
		conn, err := ln.Accept()
		if err != nil {
			break
		}
		go handleConnection(conn)
	}
	return nil
}

func onCtxDone(ctx context.Context, ln net.Listener) error {
	<-ctx.Done()
	if err := ln.Close(); err != nil {
		return err
	}
	log.Println("Listener has gracefully shut down.")
	return nil
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
