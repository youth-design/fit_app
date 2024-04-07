package app

import (
	"context"
	fit_app_auth_v0_authv0 "github.com/youth-design/fit_app/protos/gen/go/auth"
	"google.golang.org/grpc"
)

type App struct {
	fit_app_auth_v0_authv0.UnimplementedAuthServer
}

func Register(gRPCServer *grpc.Server) {
	fit_app_auth_v0_authv0.RegisterAuthServer(gRPCServer, &App{})
}

func (app *App) Register(ctx context.Context, req *fit_app_auth_v0_authv0.RegisterRequest) (*fit_app_auth_v0_authv0.RegisterResponse, error) {
	panic("implement me")
}
