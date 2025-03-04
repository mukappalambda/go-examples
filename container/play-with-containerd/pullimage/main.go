package main

import (
	"context"
	"log"

	containerd "github.com/containerd/containerd/v2/client"
	"github.com/containerd/containerd/v2/core/transfer/image"
	"github.com/containerd/containerd/v2/core/transfer/registry"
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
	ctx := namespaces.WithNamespace(context.Background(), "default")
	ref := "docker.io/library/golang:alpine"
	reg, err := registry.NewOCIRegistry(ctx, ref)
	if err != nil {
		return err
	}
	is := image.NewStore(ref)
	return client.Transfer(ctx, reg, is)
}
