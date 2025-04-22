package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	options := container.ListOptions{
		All: true,
	}
	containers, err := cli.ContainerList(context.Background(), options)
	if err != nil {
		return err
	}
	w := tabwriter.NewWriter(os.Stdout, 1, 8, 1, ' ', 0)
	defer func() {
		_ = w.Flush()
	}()
	_, _ = fmt.Fprintln(w, "CONTAINER ID\tIMAGE\tCOMMAND\tPORTS\tNAMES\t")
	for _, c := range containers {
		name := strings.TrimPrefix(c.Names[0], "/")
		fmt.Fprintf(w, "%.12s\t%.30s\t%.20q\t%+v\t%s\n", c.ID, c.Image, c.Command, c.Ports, name)
	}
	return nil
}
