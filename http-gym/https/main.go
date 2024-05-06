package main

import (
	"context"
	"flag"
	"fmt"
	"log"
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

func main() {
	flag.Parse()
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	handler := http.DefaultServeMux
	handler.HandleFunc("GET /", handleRoot())

	srv := newServer(*addr, handler)

	go func() {
		if err := srv.ListenAndServeTLS(*certFile, *keyFile); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error: %v", err)
		}
	}()
	<-ctx.Done()
	stop()
	fmt.Println("Server is shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown %s.", err)
	}
	fmt.Println("Server down.")
}

func newServer(addr string, handler http.Handler) *http.Server {
	srv := &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	return srv
}

func handleRoot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "root")
	}
}
