package main

import (
	"errors"
	httpAdapter "github.com/mukappalambda/go-examples/examples/gin-demo/internal/adapters/http"
	"log"
	"net/http"
)

func main() {
	srv := httpAdapter.NewServer()

	if err := srv.Run(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
