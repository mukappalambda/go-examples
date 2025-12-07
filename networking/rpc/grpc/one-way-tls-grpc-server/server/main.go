package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"os"

	checkinv1 "one-way-tls-grpc-server/gen/checkin/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type checkInServer struct {
	checkinv1.UnimplementedCheckInServiceServer
}

func (s *checkInServer) CheckIn(_ context.Context, req *checkinv1.CheckInRequest) (*checkinv1.CheckInResponse, error) {
	return &checkinv1.CheckInResponse{
		Message: fmt.Sprintf("response to %s", req.GetMessage()),
	}, nil
}

var _ checkinv1.CheckInServiceServer = (*checkInServer)(nil)

var (
	certFile = "server_crt.pem"
	keyFile  = "server_key.pem"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return fmt.Errorf("failed to load key pair: %s", err)
	}
	var lc net.ListenConfig
	address := "localhost:9090"
	ln, err := lc.Listen(context.Background(), "tcp", address)
	if err != nil {
		return fmt.Errorf("failed to bind address: %s", err)
	}
	defer ln.Close()
	s := grpc.NewServer(grpc.Creds(credentials.NewServerTLSFromCert(&cert)))
	checkinv1.RegisterCheckInServiceServer(s, &checkInServer{})
	if err := s.Serve(ln); err != nil {
		return fmt.Errorf("failed to serve: %s", err)
	}
	return nil
}
