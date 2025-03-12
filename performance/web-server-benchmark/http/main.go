package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

var port = flag.Int("port", 8080, "server port")

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from the http server.")
	})
	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	server := &http.Server{
		Addr:        addr,
		ReadTimeout: 5 * time.Second,
	}
	fmt.Printf("server listening at %q\n", addr)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
