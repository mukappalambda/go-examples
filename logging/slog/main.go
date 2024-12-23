package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	serverId := "My-Server"
	logger := NewLogger()
	loggingMiddleware := NewLoggingMiddleware(logger, serverId)
	http.HandleFunc("GET /", loggingMiddleware(HandleIndex()))
	if err := http.ListenAndServe(":8080", nil); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
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
