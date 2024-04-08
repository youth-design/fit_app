package server

import (
	"fmt"
	"github.com/youth-design/fit_app/libs/logger"
	"github.com/youth-design/fit_app/services/auth/internal/grpc/app"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	grpcServer *grpc.Server
	port       int
	log        *logger.Logger
}

func New(log *logger.Logger, port int) *Server {
	gRPCServer := grpc.NewServer()

	app.Register(gRPCServer, log)

	return &Server{
		grpcServer: gRPCServer,
		port:       port,
		log:        log,
	}
}

func (s *Server) Start() error {
	op := "server.Start"

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	s.log.Info().Msgf("grpc server starting at %s", l.Addr().String())
	if err = s.grpcServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
