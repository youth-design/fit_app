package server

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/youth-design/fit_app/libs/logger"
	"github.com/youth-design/fit_app/services/auth/internal/grpc/app"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

type Server struct {
	grpcServer *grpc.Server
	port       int
	log        *logger.Logger
}

func New(log *logger.Logger, port int) *Server {
	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(
			logging.PayloadReceived, logging.PayloadSent,
		),
	}

	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			log.Error().Any("panic", p).Msg("Recovered from panic")

			return status.Errorf(codes.Internal, "internal error")
		}),
	}
	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(recoveryOpts...),
		logging.UnaryServerInterceptor(InterceptorLogger(log), loggingOpts...),
	))

	app.Register(gRPCServer, log)

	return &Server{
		grpcServer: gRPCServer,
		port:       port,
		log:        log,
	}
}

func InterceptorLogger(log *logger.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		switch lvl {
		case logging.LevelError:
			log.Error().Fields(getFields(fields...)).Msg(msg)
		case logging.LevelInfo:
			log.Info().Fields(getFields(fields...)).Msg(msg)
		case logging.LevelDebug:
			log.Debug().Fields(getFields(fields...)).Msg(msg)
		default:
			log.Trace().Fields(getFields(fields...)).Msg(msg)
		}
	})
}
func getFields(args ...any) map[string]interface{} {
	if len(args)%2 == 1 {
		args = append(args, "MISSING")
	}

	fields := make(map[string]interface{})

	for i := 0; i < len(args); i += 2 {
		key, ok := args[i].(string)
		if !ok {
			continue
		}
		fields[key] = args[i+1]
	}

	return fields
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

func (s *Server) GracefulShutdown() {
	s.grpcServer.GracefulStop()
}
