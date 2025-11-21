package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	checkinv1 "one-way-tls-grpc-server/gen/checkin/v1"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var certFile = "server_crt.pem"

var message = flag.String("m", "", "request message")

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR]: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	if len(*message) == 0 {
		flag.PrintDefaults()
		return errors.New("missing -m flag")
	}
	target := "localhost:9090"
	creds, err := credentials.NewClientTLSFromFile(certFile, "localhost")
	if err != nil {
		return fmt.Errorf("failed to load cert file: %s", err)
	}
	cc, err := grpc.NewClient(target, grpc.WithTransportCredentials(creds))
	if err != nil {
		return fmt.Errorf("failed to create client connection: %s", err)
	}
	defer cc.Close()
	c := checkinv1.NewCheckInServiceClient(cc)
	resp, err := c.CheckIn(context.Background(), &checkinv1.CheckInRequest{Message: *message})
	if err != nil {
		return fmt.Errorf("failed to check in: %s", err)
	}
	fmt.Printf("response: %+v\n", resp)
	return nil
}
