package grpc_server

import (
	"log/slog"

	authTokenGen "github.com/duongnln96/blog-realworld/gen/go/auth/v1"
	"github.com/duongnln96/blog-realworld/internal/auth/infras/grpc_server"

	"github.com/duongnln96/blog-realworld/pkg/config"
)

type app struct {
	configs *config.Configs
	logger  *slog.Logger

	grpcServer grpc_server.GrpcServerAdapterI

	authTokenSvc authTokenGen.AuthTokenServiceServer
}

func NewApp(
	configs *config.Configs,
	logger *slog.Logger,

	grpcServer grpc_server.GrpcServerAdapterI,

	authTokenSvc authTokenGen.AuthTokenServiceServer,
) *app {

	return &app{
		configs:      configs,
		logger:       logger,
		grpcServer:   grpcServer,
		authTokenSvc: authTokenSvc,
	}
}

func (a *app) registerHandler() {
	authTokenGen.RegisterAuthTokenServiceServer(a.grpcServer.GetServerInstance(), a.authTokenSvc)
}

func (a *app) Serve() error {
	a.registerHandler()

	return a.grpcServer.Start()
}
