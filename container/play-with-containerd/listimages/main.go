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
	ctx := namespaces.WithNamespace(context.Background(), "default")
	imageList, err := client.ImageService().List(ctx)
	if err != nil {
		return err
	}
	for _, image := range imageList {
		fmt.Printf("%+v\n", image)
	}
	return nil
}
