package server

import (
	"errors"
	"net"

	pb "github.com/mukappalambda/go-examples/networking/rpc/grpc/note_server/api/services/notes"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/mukappalambda/go-examples/networking/rpc/grpc/note_server/cmd/noted/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/orca"
)

type Server struct {
	grpcServer *grpc.Server
}

func New() (*Server, error) {
	serverOpts := []grpc.ServerOption{
		orca.CallMetricsServerOption(nil),
		grpc.UnaryInterceptor(services.StupidInterceptor),
	}
	grpcServer := grpc.NewServer(serverOpts...)
	healthgrpc.RegisterHealthServer(grpcServer, health.NewServer())

	pb.RegisterNoteServiceServer(grpcServer, services.NewServer())
	s := &Server{
		grpcServer: grpcServer,
	}
	return s, nil
}

func (s *Server) Serve(l net.Listener) error {
	if err := s.grpcServer.Serve(l); err != nil && !errors.Is(err, net.ErrClosed) {
		return err
	}
	return nil
}
