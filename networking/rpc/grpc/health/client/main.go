package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	foov1 "github.com/mukappalambda/go-examples/grpc-health/gen/foo/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	_ "google.golang.org/grpc/health"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
)

const serviceConfig = `{
	"loadBalancingPolicy": "round_robin",
	"healthCheckConfig": {
		"serviceName": ""
	}
}`

var message = flag.String("message", "", "request message")

func main() {
	flag.Parse()
	if len(*message) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	r := manual.NewBuilderWithScheme("my-scheme")
	r.InitialState(resolver.State{
		Addresses: []resolver.Address{
			{
				Addr: "localhost:9090",
			},
		},
	})
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithResolvers(r),
		grpc.WithDefaultServiceConfig(serviceConfig),
	}
	conn, err := grpc.NewClient("127.0.0.1: 9090", options...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create client: %s", err)
		os.Exit(1)
	}
	defer conn.Close()
	fooClient := foov1.NewFooServiceClient(conn)
	ctx := context.Background()
	resp, err := fooClient.UnaryFoo(ctx, &foov1.UnaryFooRequest{
		Message: *message,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to call request: %s", err)
		os.Exit(1)
	}
	fmt.Println(resp.GetMessage())
}
