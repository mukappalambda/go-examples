package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	certFile = "server.crt"
	keyFile  = "server.key"
)

func main() {
	addr := ":9443"
	http.HandleFunc("GET /data", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("data")); err != nil {
			http.Error(w, "failed to write response", http.StatusInternalServerError)
		}
	})
	fmt.Printf("Server is listening on address: %s\n", addr)
	server := &http.Server{
		Addr:              addr,
		Handler:           nil,
		ReadHeaderTimeout: 500 * time.Millisecond,
		ReadTimeout:       500 * time.Millisecond,
		WriteTimeout:      500 * time.Millisecond,
	}
	if err := server.ListenAndServeTLS(certFile, keyFile); err != nil {
		fmt.Fprintf(os.Stderr, "failed to serve: %s\n", err)
		os.Exit(1)
	}
}
