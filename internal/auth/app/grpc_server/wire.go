//go:build wireinject
// +build wireinject

package grpc_server

import (
	"log"
	"log/slog"

	authTokenRepo "github.com/duongnln96/blog-realworld/internal/auth/adapter/repo/syclladb/auth_token"
	authTokenHandler "github.com/duongnln96/blog-realworld/internal/auth/app/grpc_server/handler/auth_token"
	authTokenUC "github.com/duongnln96/blog-realworld/internal/auth/usecases/auth_token"
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
		scylladbAdapter,
		jwtTokenAdapter,
		authTokenRepo.RepoManagerSet,
		authTokenUC.UsecasesSet,
		authTokenHandler.HandlerSet,
	))
}

func scylladbAdapter(cfg *config.Configs) (scylladb.ScyllaDBAdaterI, func()) {
	adapter := scylladb.NewScyllaDBAdapter(cfg.ScyllaDBConfigMap.Get("scylladb"))

	return adapter, func() { adapter.Close() }
}

func jwtTokenAdapter(cfg *config.Configs) token.TokenMakerI {
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
