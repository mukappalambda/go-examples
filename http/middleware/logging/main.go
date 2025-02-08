package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	handler := http.DefaultServeMux
	handler.HandleFunc("GET /", root())
	handler.Handle("GET /verbose", LoggerHandler(root()))

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           handler,
		ReadHeaderTimeout: 300 * time.Millisecond,
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func LoggerHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s - %s", r.Method, r.Host, r.Header)
		h.ServeHTTP(w, r)
	})
}

func root() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "root")
	}
}
