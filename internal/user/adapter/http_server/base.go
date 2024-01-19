package http_server

import (
	profileHandler "github.com/duongnln96/blog-realworld/internal/user/adapter/http_server/handler/profile"
	userHandler "github.com/duongnln96/blog-realworld/internal/user/adapter/http_server/handler/user"
	echoFramework "github.com/duongnln96/blog-realworld/internal/user/infras/echo_framework"

	"github.com/duongnln96/blog-realworld/pkg/config"
)

type app struct {
	config *config.Configs

	httpServer echoFramework.HTTPServerI

	userHandler   userHandler.HandlerI
	profileHander profileHandler.HandlerI
}

func NewApp(
	config *config.Configs,
	httpServer echoFramework.HTTPServerI,

	userHandler userHandler.HandlerI,
	profileHander profileHandler.HandlerI,
) *app {
	return &app{
		config:     config,
		httpServer: httpServer,

		userHandler:   userHandler,
		profileHander: profileHander,
	}
}

func (a *app) Serve() error {
	a.routeVer1()

	return a.httpServer.Start()
}
