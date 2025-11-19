package main

import (
	"context"
	"fmt"
	"os"
	"time"

	calcv1 "github.com/mukappalambda/grpc/headerandtrailer/gen/calc/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	target := "localhost:9090"
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(target, options...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create grpc client: %s", err)
		return
	}
	defer conn.Close()
	client := calcv1.NewCalcServiceClient(conn)

	callUnaryCalc(client)
	callStreamCalc(client)
}

func callUnaryCalc(c calcv1.CalcServiceClient) {
	var header metadata.MD
	var trailer metadata.MD
	resp, err := c.UnaryCalc(context.Background(), &calcv1.UnaryCalcRequest{Num: 123}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to unary calc: %s", err)
		return
	}
	fmt.Printf("%+v\n", resp)
	vs := header.Get("my-header-key")
	for _, v := range vs {
		fmt.Printf("my-header-key: %s\n", v)
	}
	if ds, ok := trailer["datetime"]; ok {
		for _, d := range ds {
			dt, err := time.Parse(time.RFC3339, d)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
				return
			}
			fmt.Printf("datetime: %s", dt.UTC())
		}
	}
}

func callStreamCalc(c calcv1.CalcServiceClient) {
	stream, err := c.StreamCalc(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create stream: %s\n", err)
	}
	for i := range 5 {
		m := &calcv1.StreamCalcRequest{
			Num: float32(i),
		}
		if err := stream.Send(m); err != nil {
			fmt.Fprintf(os.Stderr, "failed to send message: %s\n", err)
			return
		}
	}
	m, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to close and receive: %s", err)
	}
	fmt.Println(m.GetNum())
}
