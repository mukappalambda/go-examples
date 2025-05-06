package main

import (
	"fmt"
	"os"

	"github.com/mukappalambda/go-examples/networking/rpc/kvstore/server"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	addr := "127.0.0.1:8080"
	remoteStoreService := server.NewRemoteStoreService()
	if err := remoteStoreService.Run(addr); err != nil {
		return err
	}
	return nil
}
