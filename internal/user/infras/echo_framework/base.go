package echo_framework

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/duongnln96/blog-realworld/internal/pkg/serror"
	"github.com/duongnln96/blog-realworld/internal/pkg/validator"
	"github.com/duongnln96/blog-realworld/pkg/config"
	userOpenApi3 "github.com/duongnln96/blog-realworld/third_party/OpenAPI/user"
)

//go:embed assets
var content embed.FS

var _ (HTTPServerI) = (*httpServer)(nil)

var HTTPServerSet = wire.NewSet(NewHttpServer)

type HTTPServerI interface {
	GroupRouter(groupPrefix string, m ...echo.MiddlewareFunc) *echo.Group
	Start() error
	Stop() error
}

type httpServer struct {
	configs *config.Configs

	e *echo.Echo
}

func NewHttpServer() HTTPServerI {

	var echoServer = echo.New()
	echoServer.HTTPErrorHandler = serror.CustomEchoErrorHandler
	echoServer.Validator = validator.NewSValidator()

	return &httpServer{
		e: echoServer,
	}
}

func (s *httpServer) initSystemRouter() {

	// health route
	s.e.GET("/health", serror.HealthHandler)

	// swagger
	fsys, _ := fs.Sub(content, "assets")
	s.e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", http.FileServer(http.FS(fsys)))))

	// register openapi3
	s.registerOpenApi3()
}

func (s *httpServer) registerOpenApi3() {
	swagger := userOpenApi3.NewOpenAPI3()
	s.e.GET("/user/swagger.json", func(c echo.Context) error {
		return c.JSON(http.StatusOK, swagger)
	})
}

func (s *httpServer) initSystemMiddleware() {
	s.e.Use(middleware.RequestID())
	s.e.Use(middleware.Recover())
	s.e.Use(middleware.Secure())
	// s.e.Use(middleware.LoggerWithConfig())

	s.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost", "*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))
}

func (s *httpServer) Start() error {

	serverAddress := fmt.Sprintf("%s:%d", s.configs.Other.Get("http_address"), s.configs.Other.Get("http_port"))

	s.initSystemMiddleware()
	s.initSystemRouter()

	go func() {
		slog.Info("🌏 Start server...")
		if err := s.e.Start(serverAddress); errors.Is(err, http.ErrServerClosed) {
			slog.Info("=> shutting down the server", "error_info", err.Error())
			return
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.e.Shutdown(ctx); err != nil {
		slog.Info("=> Gracefully shutting down the server", "err_info", err.Error())
		return err
	}

	return nil
}

func (s *httpServer) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.e.Shutdown(ctx); err != nil {
		slog.Info("=> Gracefully shutting down the server", "err_info", err.Error())
	}

	return nil
}

func (s *httpServer) GroupRouter(groupPrefix string, m ...echo.MiddlewareFunc) *echo.Group {
	return s.e.Group(groupPrefix, m...)
}
