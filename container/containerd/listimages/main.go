package main

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	containerd "github.com/containerd/containerd/v2/client"
	"github.com/containerd/containerd/v2/pkg/namespaces"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
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
	tw := tabwriter.NewWriter(os.Stdout, 1, 8, 1, ' ', 0)
	fmt.Fprintf(tw, "REF\tTYPE\tDIGEST\n")
	for _, image := range imageList {
		fmt.Fprintf(tw, "%v\t%v\t%v\n", image.Name, image.Target.MediaType, image.Target.Digest)
	}
	return tw.Flush()
}
