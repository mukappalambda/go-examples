package main

import (
	"context"
	"fmt"
	"log"

	containerd "github.com/containerd/containerd/v2/client"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		return err
	}
	defer client.Close()
	serviceInfo, err := client.IntrospectionService().Server(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", serviceInfo)
	return nil
}
