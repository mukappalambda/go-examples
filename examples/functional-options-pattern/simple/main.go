package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var addr = flag.String("addr", ":8080", "server address")

func main() {
	flag.Parse()

	srv := New(
		WithAddr(*addr),
		WithReadTimeout(500*time.Millisecond),
		WithReadHeaderTimeout(500*time.Millisecond),
	)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

type options struct {
	addr              string
	handler           http.Handler
	readTimeout       time.Duration
	readHeaderTimeout time.Duration
}

type Option func(*options)

func WithAddr(addr string) Option {
	return func(options *options) {
		options.addr = addr
	}
}

func WithHandler(mux *http.ServeMux) Option {
	return func(options *options) {
		var handler http.Handler = mux
		options.handler = handler
	}
}

func WithReadTimeout(timeout time.Duration) Option {
	return func(options *options) {
		options.readTimeout = timeout
	}
}

func WithReadHeaderTimeout(timeout time.Duration) Option {
	return func(options *options) {
		options.readHeaderTimeout = timeout
	}
}

func addRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /{$}", handleIndex())
	mux.HandleFunc("GET /greet/{name}", handleGreet("Hiii"))
}

func New(opts ...Option) *http.Server {
	mux := http.DefaultServeMux
	addRoutes(mux)
	var options options
	for _, opt := range opts {
		opt(&options)
	}

	srv := &http.Server{
		Addr:              options.addr,
		Handler:           options.handler,
		ReadTimeout:       options.readTimeout,
		ReadHeaderTimeout: options.readHeaderTimeout,
	}
	return srv
}

func handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s\n", "index route")
	}
}

func handleGreet(prefix string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		fmt.Fprintf(w, "%s, %s\n", prefix, name)
	}
}
