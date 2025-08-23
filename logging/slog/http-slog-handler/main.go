package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	defaultReadTimeout       = 500 * time.Millisecond
	defaultReadHeaderTimeout = 500 * time.Millisecond

	addr              = flag.String("addr", ":8080", "server address")
	readTimeout       = flag.Duration("readtimeout", defaultReadTimeout, "server readtimeout")
	readHeaderTimeout = flag.Duration("readheadertimeout", defaultReadHeaderTimeout, "server readheadertimeout")
)

type Option func(*http.Server)

func WithReadTimeout(timeout time.Duration) Option {
	return func(s *http.Server) { s.ReadTimeout = timeout }
}

func WithReadHeaderTimeout(timeout time.Duration) Option {
	return func(s *http.Server) { s.ReadHeaderTimeout = timeout }
}

func main() {
	flag.Parse()

	var opts []Option
	opts = append(opts, WithReadTimeout(*readTimeout))
	opts = append(opts, WithReadHeaderTimeout(*readHeaderTimeout))

	server := NewServer(*addr, opts...)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}()
	fmt.Printf("Server is listening on %s\n", server.Addr)
	wg.Wait()
}

func NewServer(addr string, opts ...Option) *http.Server {
	handler := newHandler()
	server := &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadTimeout:       defaultReadTimeout,
		ReadHeaderTimeout: defaultReadHeaderTimeout,
	}
	for _, o := range opts {
		o(server)
	}
	return server
}

func newHandler() http.Handler {
	logger := NewLogger()
	serverId := "My-Server"
	loggingMiddleware := NewLoggingMiddleware(logger, serverId)
	mux := http.DefaultServeMux
	mux.HandleFunc("GET /", loggingMiddleware(HandleIndex()))
	return mux
}

func NewLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return logger
}

func NewLoggingMiddleware(logger *slog.Logger, msg string) func(f http.HandlerFunc) http.HandlerFunc {
	loggingMiddleWare := func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			f.ServeHTTP(w, r)
			logger.Info(msg, "Method", r.Method, "Path", r.URL.Path, "Query-String", r.URL.Query(), "User-Agent", r.UserAgent(), "Proto", r.Proto)
		}
	}
	return loggingMiddleWare
}

func HandleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "index page")
	}
}
