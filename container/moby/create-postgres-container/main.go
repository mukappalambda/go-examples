package main

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	clientOpts := []client.Opt{
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	}
	cli, err := client.NewClientWithOpts(clientOpts...)
	if err != nil {
		return fmt.Errorf("failed to new a client: %s", err)
	}
	cfgEnv := []string{"POSTGRES_USER=postgres", "POSTGRES_PASSWORD=password", "POSTGRES_DB=demo"}
	ctx := context.Background()
	config := &container.Config{
		Image: "postgres:16",
		Env:   cfgEnv,
		ExposedPorts: nat.PortSet{
			"5432/tcp": {},
		},
	}
	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"5432/tcp": {
				{HostIP: "0.0.0.0", HostPort: "5432"},
			},
		}}
	containerName := "postgres"
	if err := startPostgresContainer(ctx, cli, config, hostConfig, containerName); err != nil {
		return err
	}
	return nil
}

func startPostgresContainer(ctx context.Context, cli *client.Client, config *container.Config, hostConfig *container.HostConfig, containerName string) error {
	res, err := cli.ContainerCreate(ctx, config, hostConfig, nil, nil, containerName)
	if err != nil {
		return fmt.Errorf("failed to create the container: %s", err)
	}
	if err := cli.ContainerStart(ctx, res.ID, container.StartOptions{}); err != nil {
		return fmt.Errorf("failed to start the container: %s", err)
	}
	fmt.Printf("container %q is created", res.ID)
	return nil
}
