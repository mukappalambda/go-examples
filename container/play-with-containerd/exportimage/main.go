package main

import (
	"context"
	"fmt"
	"log"
	"os"

	containerd "github.com/containerd/containerd/v2/client"
	tarchive "github.com/containerd/containerd/v2/core/transfer/archive"
	"github.com/containerd/containerd/v2/core/transfer/image"
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
		return fmt.Errorf("failed to connect to the containerd instance: %w", err)
	}
	defer client.Close()
	ctx := namespaces.WithNamespace(context.Background(), "default")
	name := "docker.io/library/golang:alpine"
	out := "out.tar"
	w, err := os.Create(out)
	if err != nil {
		return fmt.Errorf("failed to create %q", name)
	}
	defer w.Close()
	exportOpts := []tarchive.ExportOpt{}
	storeOpts := []image.StoreOpt{image.WithExtraReference(name)}
	src := image.NewStore("", storeOpts...)
	dest := tarchive.NewImageExportStream(w, "", exportOpts...)

	if err := client.Transfer(ctx, src, dest); err != nil {
		return fmt.Errorf("failed to export images: %w", err)
	}
	fmt.Printf("%q is exported to %q\n", name, out)
	return nil
}
