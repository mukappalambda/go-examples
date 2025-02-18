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

type Server struct {
	*http.Server
}

var (
	port            = flag.Uint("port", 8080, "server port")
	readTimeout     = flag.Duration("read-timeout", 500*time.Millisecond, "server read timeout")
	shutdownTimeout = flag.Duration("shutdown-timeout", 5*time.Second, "server shutdown timeout")
)

func main() {
	flag.Parse()
	addr := fmt.Sprintf(":%d", *port)
	srv := newServer(addr, *readTimeout)
	srv.setupRoutes()
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.run(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	stop()
	fmt.Println("Server is gracefully shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), *shutdownTimeout)
	if err := srv.Shutdown(ctx); err != nil {
		cancel()
		log.Fatalf("Server forces to shut down. %s", err)
	}
	cancel()
	fmt.Println("Server is down.")
}

func newServer(addr string, readTimeout time.Duration) *Server {
	return &Server{
		&http.Server{Addr: addr, ReadTimeout: readTimeout},
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
