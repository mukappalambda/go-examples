package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	*http.Server
}

func main() {
	srv := newServer()
	srv.setupRoutes()
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := srv.run(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	stop()
	fmt.Println("Server is gracefully shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forces to shut down. %s", err)
	}
	fmt.Println("Server is down.")
}

func newServer() *Server {
	return &Server{
		&http.Server{
			Addr: ":8080",
		},
	}
}

func (srv *Server) setupRoutes() {
	handler := http.NewServeMux()
	handler.HandleFunc("GET /", handleRoot())
	handler.HandleFunc("GET /hello", handleGreet("hello"))
	handler.HandleFunc("GET /hiii", handleGreet("hiii"))
	handler.HandleFunc("GET /slow", handleSlow(3*time.Second))
	srv.Handler = handler
}

func (srv *Server) run() error {
	return srv.ListenAndServe()
}

func handleRoot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "root")
	}
}

func handleGreet(msg string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, msg)
	}
}
func handleSlow(duration time.Duration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		fmt.Fprint(w, "ok")
	}
}
