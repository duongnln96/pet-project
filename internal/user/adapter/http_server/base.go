package http_server

import (
	profileHandler "github.com/duongnln96/blog-realworld/internal/user/adapter/http_server/handler/profile"
	userHandler "github.com/duongnln96/blog-realworld/internal/user/adapter/http_server/handler/user"
	echoFramework "github.com/duongnln96/blog-realworld/internal/user/infras/echo_framework"

	"github.com/duongnln96/blog-realworld/pkg/config"
)

func Run(config *config.Configs) error {
	app, cancel := InitNewApp(config)
	defer cancel()

	return app.httpServer.Start()
}

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
