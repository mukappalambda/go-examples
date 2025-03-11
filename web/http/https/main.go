package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	addr     = flag.String("addr", ":4443", "addr")
	certFile = flag.String("certFile", "server.crt", "certificate file")
	keyFile  = flag.String("keyFile", "server.key", "private key file")
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	flag.Parse()
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	handler := http.DefaultServeMux
	handler.HandleFunc("GET /", handleRoot())

	srv := newServer(*addr, handler)

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

func newServer(addr string, handler http.Handler) *http.Server {
	srv := &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadHeaderTimeout: 300 * time.Millisecond,
	}
	return srv
}

func handleRoot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "root")
	}
}
