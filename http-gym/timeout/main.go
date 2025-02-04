package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	h := handler(500 * time.Millisecond)
	handlerWithTimeout := http.TimeoutHandler(h, 300*time.Millisecond, "timeout is reached")

	// curl http://localhost:8080/notimeout => 200 OK
	http.HandleFunc("GET /notimeout", h)
	// curl http://localhost:8080/withtimeout => 503 Service Unavailable
	http.Handle("GET /withtimeout", handlerWithTimeout)
	server := &http.Server{Addr: ":8080", ReadHeaderTimeout: 300 * time.Millisecond}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func handler(d time.Duration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(d)
		fmt.Fprint(w, "root")
	}
}
