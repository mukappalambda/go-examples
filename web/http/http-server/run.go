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

	"github.com/mukappalambda/go-examples/web/http/http-server/handlers"
	"github.com/mukappalambda/go-examples/web/http/http-server/server"
)

var (
	port              = flag.Uint("port", 8080, "server port")
	readTimeout       = flag.Duration("read-timeout", 500*time.Millisecond, "server read timeout")
	readHeaderTimeout = flag.Duration("read-header-timeout", 500*time.Millisecond, "server read header timeout")
	writeTimeout      = flag.Duration("write-timeout", 500*time.Millisecond, "server write timeout")
	shutdownTimeout   = flag.Duration("shutdown-timeout", 5*time.Second, "server shutdown timeout")
)

func run() error {
	flag.Parse()
	addr := fmt.Sprintf(":%d", *port)
	sopts := []server.Opt{
		server.WithDefaultReadTimeout(*readTimeout),
		server.WithDefaultReadHeaderTimeout(*readHeaderTimeout),
		server.WithDefaultWriteTimeout(*writeTimeout),
		server.WithDefaultHandler(handlers.DefaultMux()),
	}
	server := server.New(addr, sopts...)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("Server is running at %q\n", addr)
	go func() {
		if err := server.Run(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	stop()
	fmt.Println("Server is gracefully shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), *shutdownTimeout)
	if err := server.Shutdown(ctx); err != nil {
		cancel()
		return fmt.Errorf("server forces to shut down. %s", err)
	}
	defer cancel()
	fmt.Println("Server is down.")
	return nil
}
