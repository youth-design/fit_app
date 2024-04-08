package app

import (
	"context"
	"github.com/youth-design/fit_app/libs/logger"
	fit_app_auth_v0_authv0 "github.com/youth-design/fit_app/protos/gen/go/auth"
	"google.golang.org/grpc"
)

type App struct {
	fit_app_auth_v0_authv0.AuthServer
	log         *logger.Logger
	Description *grpc.ServiceDesc
}

func New(log *logger.Logger) *App {
	return &App{
		log:         log,
		Description: &fit_app_auth_v0_authv0.Auth_ServiceDesc,
	}
}

func (app *App) Register(ctx context.Context, req *fit_app_auth_v0_authv0.RegisterRequest) (*fit_app_auth_v0_authv0.RegisterResponse, error) {
	panic("implement me")
}
