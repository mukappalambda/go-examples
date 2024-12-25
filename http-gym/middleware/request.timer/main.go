package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	addr := ":8080"
	http.Handle("GET /data", requestTimer(handleData()))
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := http.ListenAndServe(addr, nil); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error serving on %s\n", err)
		}
	}()
	fmt.Printf("server running on %q\n", addr)
	wg.Wait()
}

func requestTimer(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		elapsed := time.Since(start)
		log.Printf("[addr] %s  [method] %s  [path] %s  [elapsed time]: %s\n", r.RemoteAddr, r.Method, r.URL.Path, elapsed.String())
	})
}

func handleData() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "data")
		})
}
