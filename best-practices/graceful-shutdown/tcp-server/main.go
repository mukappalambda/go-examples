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
	"sync"
	"syscall"
	"time"
)

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
	var lc net.ListenConfig
	ln, err := lc.Listen(ctx, "tcp", "127.0.0.1:8080")
	if err != nil {
		return fmt.Errorf("error listening on the network address: %w", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func(ln net.Listener) {
		defer wg.Done()
		for {
			conn, err := ln.Accept()
			if err != nil {
				break
			}
			go handleConnection(conn)
		}
	}(ln)
	log.Printf("Server is listening on %q\n", ln.Addr())
	<-ctx.Done()
	log.Println("Server is shutting down...")
	timeout := 100 * time.Millisecond
	shutdownCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		<-ctx.Done()
		if e := ln.Close(); e != nil {
			err = e
		}
		log.Println("Listener has gracefully shut down.")
	}(shutdownCtx)
	wg.Wait()
	if err != nil {
		return fmt.Errorf("error on shutting down the server: %s", err)
	}
	log.Println("Server shut down successfully.")
	return nil
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
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
