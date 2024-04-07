package app

import (
	"context"
	"github.com/youth-design/fit_app/libs/logger"
	fit_app_auth_v0_authv0 "github.com/youth-design/fit_app/protos/gen/go/auth"
	"google.golang.org/grpc"
)

type App struct {
	fit_app_auth_v0_authv0.UnimplementedAuthServer
	log *logger.Logger
}

func Register(gRPCServer *grpc.Server, log *logger.Logger) {
	fit_app_auth_v0_authv0.RegisterAuthServer(
		gRPCServer,
		&App{
			log: log,
		},
	)
}

func (app *App) Register(ctx context.Context, req *fit_app_auth_v0_authv0.RegisterRequest) (*fit_app_auth_v0_authv0.RegisterResponse, error) {
	panic("implement me")
}
