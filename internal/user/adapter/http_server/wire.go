//go:build wireinject
// +build wireinject

package http_server

import (
	"github.com/google/wire"

	"github.com/duongnln96/blog-realworld/pkg/adapter/postgres"
	"github.com/duongnln96/blog-realworld/pkg/config"

	profileHandler "github.com/duongnln96/blog-realworld/internal/user/adapter/http_server/handler/profile"
	userHandler "github.com/duongnln96/blog-realworld/internal/user/adapter/http_server/handler/user"
	"github.com/duongnln96/blog-realworld/internal/user/adapter/http_server/middlewares"
	followRepo "github.com/duongnln96/blog-realworld/internal/user/adapter/repository/postgresql/follow"
	userRepo "github.com/duongnln96/blog-realworld/internal/user/adapter/repository/postgresql/user"

	echoAdapaer "github.com/duongnln96/blog-realworld/internal/user/infras/echo_framework"
	authTokenDomainClient "github.com/duongnln96/blog-realworld/internal/user/infras/grpc_client/auth/auth_token"

	"github.com/duongnln96/blog-realworld/internal/user/core/port"
	profileService "github.com/duongnln96/blog-realworld/internal/user/core/service/profile"
	userService "github.com/duongnln96/blog-realworld/internal/user/core/service/user"
)

func InitNewApp(
	config *config.Configs,
) (*app, func(), error) {
	panic(wire.Build(
		NewApp,
		newHTTPServer,

		newAuthTokenDoamin,
		middlewares.AuthMiddlewareSet,

		newPostgresDbAdapter,

		userRepo.RepositorySet,
		userService.ServiceSet,
		userHandler.HandlerSet,

		followRepo.RepositorySet,
		profileService.ServiceSet,
		profileHandler.HandlerSet,
	))
}

func newPostgresDbAdapter(cfg *config.Configs) (postgres.PostgresDBAdapterI, func()) {
	adapter := postgres.NewPostgresDBAdapter(cfg.PostgresConfigMap.Get("postgres"))

	return adapter, func() { adapter.Close() }
}

func newHTTPServer() (echoAdapaer.HTTPServerI, func()) {
	echoServer := echoAdapaer.NewHttpServer()

	return echoServer, func() { echoServer.Stop() }
}

func newAuthTokenDoamin(cfg *config.Configs) port.AuthTokenDomainI {
	return authTokenDomainClient.NewAuthTokenDomainClient(cfg)
}
