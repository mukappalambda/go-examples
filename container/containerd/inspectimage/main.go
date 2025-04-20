package main

import (
	"context"
	"fmt"
	"log"

	containerd "github.com/containerd/containerd/v2/client"
	"github.com/containerd/containerd/v2/pkg/namespaces"
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
	imageStore := client.ImageService()
	ctx := namespaces.WithNamespace(context.Background(), "default")
	img, err := imageStore.Get(ctx, "docker.io/library/golang:alpine")
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", img)
	return nil
}
