package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	srv := newServer()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	var err error
	go func() {
		err = srv.ListenAndServe()
	}()

	if err != nil && err != http.ErrServerClosed {
		return err
	}
	fmt.Printf("Server is running at %q\n", srv.Addr)
	<-c
	fmt.Println("Gracefully shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("timeout is reached: %w", err)
	}
	fmt.Println("Server has been shut down.")
	return nil
}

func newServer() *http.Server {
	router := newRouter()
	setupRoutes(router)
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadHeaderTimeout: 300 * time.Millisecond,
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
