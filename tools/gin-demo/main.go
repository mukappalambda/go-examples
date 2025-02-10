package main

import (
	"errors"
	"log"
	"net/http"

	httpAdapter "github.com/mukappalambda/go-examples/gin-demo/internal/adapters/http"
)

func main() {
	srv := httpAdapter.NewServer()

	if err := srv.Run(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
