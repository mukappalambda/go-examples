package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	calcv1 "github.com/mukappalambda/grpc/headerandtrailer/gen/calc/v1"
	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type calcServer struct {
	calcv1.UnimplementedCalcServiceServer
}

func (s *calcServer) UnaryCalc(ctx context.Context, req *calcv1.UnaryCalcRequest) (*calcv1.UnaryCalcResponse, error) {
	var mdErr error
	defer func() {
		header := metadata.Pairs("my-header-key", "my-header-value")
		if err := grpc.SetHeader(ctx, header); err != nil {
			fmt.Fprintf(os.Stderr, "failed to set header: %s", err)
			mdErr = err
		}
		trailer := metadata.Pairs("datetime", time.Now().In(time.UTC).Format(time.RFC3339))
		if err := grpc.SetTrailer(ctx, trailer); err != nil {
			fmt.Fprintf(os.Stderr, "failed to set trailer: %s", err)
			mdErr = err
		}
	}()
	num := req.GetNum()
	if num < 0 {
		st := status.New(codes.InvalidArgument, "invalid num")
		ds, err := st.WithDetails(&epb.BadRequest_FieldViolation{
			Field:       "num",
			Description: "invalid num",
			Reason:      "Make the num a positive float32",
		})
		if err != nil {
			return nil, st.Err()
		}
		return nil, ds.Err()
	}
	return &calcv1.UnaryCalcResponse{Num: num}, mdErr
}

func (s *calcServer) StreamCalc(stream grpc.ClientStreamingServer[calcv1.StreamCalcRequest, calcv1.StreamCalcResponse]) error {
	var total float32
	total = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		total += req.GetNum()
	}
	return stream.SendAndClose(&calcv1.StreamCalcResponse{Num: total})
}

var _ calcv1.CalcServiceServer = (*calcServer)(nil)

func main() {
	network := "tcp"
	address := "localhost:9090"
	var lc net.ListenConfig
	ctx := context.Background()
	ln, err := lc.Listen(ctx, network, address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to listen address: %s\n", err)
		return
	}
	server := grpc.NewServer()
	calcv1.RegisterCalcServiceServer(server, &calcServer{})
	if err := server.Serve(ln); err != nil {
		fmt.Fprintf(os.Stderr, "failed to serve grpc server: %s\n", err)
		return
	}
}
