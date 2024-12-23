package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	srv := newServer()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-c
	fmt.Println("Gracefully shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Timeout is reached: %s", err)
	}
	fmt.Println("Server has been shut down.")
}

func newServer() *http.Server {
	router := newRouter()
	setupRoutes(router)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	return srv
}

func newRouter() *http.ServeMux {
	return http.NewServeMux()
}

func setupRoutes(router *http.ServeMux) {
	router.HandleFunc("/", handleRoot())

}

func handleRoot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		fmt.Fprintln(w, "root")
	}
}
