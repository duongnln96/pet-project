package http_server

import (
	"context"
	"embed"
	"errors"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/duongnln96/blog-realworld/internal/pkg/validator"
	profileHandler "github.com/duongnln96/blog-realworld/internal/user/app/http_server/handler/profile"
	userHandler "github.com/duongnln96/blog-realworld/internal/user/app/http_server/handler/user"
	"github.com/duongnln96/blog-realworld/internal/user/app/open_api3"

	"github.com/duongnln96/blog-realworld/pkg/config"
)

//go:embed assets
var content embed.FS

func Run(config *config.Configs) error {
	httpApp := InitNewApp(config)

	return httpApp.Serve()
}

type app struct {
	config *config.Configs

	userHandler   userHandler.HandlerI
	profileHander profileHandler.HandlerI
}

func NewApp(
	config *config.Configs,

	userHandler userHandler.HandlerI,
	profileHander profileHandler.HandlerI,
) *app {
	return &app{
		config:        config,
		userHandler:   userHandler,
		profileHander: profileHander,
	}
}

func (s *app) initRouter(e *echo.Echo) {

	// health route
	e.GET("/health", serror.HealthHandler)

	// swagger
	fsys, _ := fs.Sub(content, "assets")
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", http.FileServer(http.FS(fsys)))))

	// register openapi3
	s.registerOpenApi3(e)

	s.routeVer1(e)
}

func (s *app) registerOpenApi3(e *echo.Echo) {
	swagger := open_api3.NewOpenAPI3()
	e.GET("/user/swagger.json", func(c echo.Context) error {
		return c.JSON(http.StatusOK, swagger)
	})
}

func (s *app) initMiddleware(e *echo.Echo) {
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	// e.Use(middleware.LoggerWithConfig())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost", "*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut},
	}))
}

func (s *app) Serve() error {

	var e = echo.New()
	e.HTTPErrorHandler = serror.CustomEchoErrorHandler
	e.Validator = validator.NewSValidator()

	s.initMiddleware(e)
	s.initRouter(e)

	go func() {
		slog.Info("🌏 Start server...")
		if err := e.Start(":80"); errors.Is(err, http.ErrServerClosed) {
			slog.Info("=> shutting down the server", "error_info", err.Error())
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		slog.Info("=> Gracefully shutting down the server", "err_info", err.Error())
	}

	return nil
}
