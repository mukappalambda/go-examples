package main

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return fmt.Errorf("failed to new client opts: %s", err)
	}
	// docker images
	images, err := cli.ImageList(context.Background(), image.ListOptions{
		All: true,
	})
	if err != nil {
		return fmt.Errorf("failed to list images: %s", err)
	}

	for _, image := range images {
		fmt.Printf("%q\t%q\n", image.RepoTags, image.RepoDigests)
	}
	return nil
}
