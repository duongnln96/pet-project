//go:build wireinject
// +build wireinject

package http_server

import (
	"github.com/google/wire"

	"github.com/duongnln96/blog-realworld/pkg/adapter/postgres"
	"github.com/duongnln96/blog-realworld/pkg/config"

	profileHandler "github.com/duongnln96/blog-realworld/internal/user/app/http_server/handler/profile"
	userHandler "github.com/duongnln96/blog-realworld/internal/user/app/http_server/handler/user"
	followRepo "github.com/duongnln96/blog-realworld/internal/user/infras/repo/follow"
	userRepo "github.com/duongnln96/blog-realworld/internal/user/infras/repo/user"

	profileService "github.com/duongnln96/blog-realworld/internal/user/core/service/profile"
	userService "github.com/duongnln96/blog-realworld/internal/user/core/service/user"
)

func InitNewApp(
	config *config.Configs,
) *app {
	panic(wire.Build(
		NewApp,
		postgresDbAdapter,
		userRepo.RepositorySet,
		userService.ServiceSet,
		userHandler.HandlerSet,
		followRepo.RepositorySet,
		profileService.ServiceSet,
		profileHandler.HandlerSet,
	))
}

func postgresDbAdapter(cfg *config.Configs) postgres.PostgresDBAdapterI {
	adapter := postgres.NewPostgresDBAdapter(cfg.PostgresConfigMap.Get("postgres"))

	return adapter
}
