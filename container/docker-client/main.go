package main

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		log.Fatal(err)
	}
	// docker images
	images, err := cli.ImageList(context.Background(), image.ListOptions{
		All: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, image := range images {
		fmt.Printf("%q\t%q\n", image.RepoTags, image.RepoDigests)
	}

	// docker pull

	// docker run ...

	// docker ps -a
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{
		All: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, container := range containers {
		fmt.Printf(
			"%s\t%v\t%q\t%+q\t%v\t%s\n",
			container.ID[:12],
			container.Image,
			container.Command,
			container.Status,
			container.Ports,
			container.Names[0],
		)
	}
}
