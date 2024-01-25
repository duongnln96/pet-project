//go:build wireinject
// +build wireinject

package grpc_server

import (
	"log"
	"log/slog"

	authTokenHandler "github.com/duongnln96/blog-realworld/internal/auth/adapter/grpc_server/handler/auth_token"
	authTokenRepo "github.com/duongnln96/blog-realworld/internal/auth/adapter/repo/syclladb/auth_token"
	authTokenSvc "github.com/duongnln96/blog-realworld/internal/auth/core/service/auth_token"
	grpcServerAdapter "github.com/duongnln96/blog-realworld/internal/auth/infras/grpc_server"
	grpcMiddlewares "github.com/duongnln96/blog-realworld/pkg/middleware/grpc"
	"google.golang.org/grpc"

	"github.com/duongnln96/blog-realworld/internal/pkg/token"
	"github.com/duongnln96/blog-realworld/pkg/adapter/scylladb"
	"github.com/duongnln96/blog-realworld/pkg/config"
	"github.com/google/wire"
)

func InitNewApp(
	config *config.Configs,
	logger *slog.Logger,
) (*app, func()) {
	panic(wire.Build(
		NewApp,
		NewGrpcServerAdapter,
		NewScylladbAdapter,
		NewJwtTokenAdapter,
		authTokenRepo.RepoManagerSet,
		authTokenSvc.ServiceSet,
		authTokenHandler.HandlerSet,
	))
}

func NewScylladbAdapter(cfg *config.Configs) (scylladb.ScyllaDBAdaterI, func()) {
	adapter := scylladb.NewScyllaDBAdapter(cfg.ScyllaDBConfigMap.Get("scylladb"))

	return adapter, func() { adapter.Close() }
}

func NewJwtTokenAdapter(cfg *config.Configs) token.TokenMakerI {
	secret, ok := cfg.Other.Get("jwt_secret_key").(string)
	if !ok {
		log.Panic("Cannot get jwt_secret_key")
	}

	tokenMaker, err := token.NewJWTTokenMaker(secret)
	if err != nil {
		log.Panic("Cannot create new instance for token maker")
	}

	return tokenMaker
}

func NewGrpcServerAdapter(cfg *config.Configs, logger *slog.Logger) (grpcServerAdapter.GrpcServerAdapterI, func()) {

	adapter := grpcServerAdapter.NewGRPCServerAdapter(cfg, logger,
		grpc.ChainUnaryInterceptor(
			grpcMiddlewares.NewLoggingUnaryInterceptor(logger).LoggingUnaryInterceptor,
		),
	)

	return adapter, func() { adapter.Stop() }
}
