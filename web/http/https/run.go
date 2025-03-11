package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

var (
	addr     = flag.String("addr", ":4443", "addr")
	certFile = flag.String("certFile", "server.crt", "certificate file")
	keyFile  = flag.String("keyFile", "server.key", "private key file")
)

func run() error {
	flag.Parse()
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	opts := []Opt{
		WithDefaultHandler(defaultMux()),
		WithDefaultReadHeaderTimeout(5 * time.Second),
	}
	srv := NewServer(*addr, opts...)

	var err error
	go func() {
		err = srv.ListenAndServeTLS(*certFile, *keyFile)
	}()
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to serve HTTP server: %w", err)
	}
	fmt.Printf("Server is running at %q\n", *addr)
	<-ctx.Done()
	stop()
	fmt.Println("Server is shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("server forced to shutdown %w", err)
	}
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("server shut down ungracefully: %w", err)
	}
	fmt.Println("Server down.")
	return nil
}
