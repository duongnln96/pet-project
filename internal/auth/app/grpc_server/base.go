package grpc_server

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	authTokenGen "github.com/duongnln96/blog-realworld/gen/go/auth/v1"

	"github.com/duongnln96/blog-realworld/pkg/logger"

	grpcMiddlewares "github.com/duongnln96/blog-realworld/pkg/middleware/grpc"

	"github.com/duongnln96/blog-realworld/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/sirupsen/logrus"
)

func RunGRPCServer(configs *config.Configs) error {

	// setup logger
	logrus.SetLevel(logger.ConvertLogLevel("debug"))
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)

	logger := slog.New(logger.NewLogrusHandler(logrus.StandardLogger()))

	app, cancel := InitNewApp(configs, logger)
	defer cancel()

	return app.Serve()
}

type app struct {
	configs *config.Configs
	logger  *slog.Logger

	authTokenSvc authTokenGen.AuthTokenServiceServer
}

func NewApp(
	configs *config.Configs,
	logger *slog.Logger,

	authTokenSvc authTokenGen.AuthTokenServiceServer,
) *app {

	return &app{
		configs:      configs,
		logger:       logger,
		authTokenSvc: authTokenSvc,
	}
}

func (a *app) registerHandler(grpcs *grpc.Server) {
	authTokenGen.RegisterAuthTokenServiceServer(grpcs, a.authTokenSvc)

}

func (a *app) registerReflection(grpcs *grpc.Server) {
	reflection.Register(grpcs)
}

func (a *app) initUnaryInterceptor() grpc.ServerOption {

	loggingInter := grpcMiddlewares.NewLoggingUnaryInterceptor(a.logger)

	unaryOptions := grpc.ChainUnaryInterceptor(loggingInter.LoggingUnaryInterceptor)

	return unaryOptions
}

func (a *app) Serve() error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server := grpc.NewServer(
		a.initUnaryInterceptor(),
	)

	go func() {
		defer server.GracefulStop()
		<-ctx.Done()
	}()

	// register handler
	a.registerHandler(server)

	// register reflection
	a.registerReflection(server)

	// gRPC Server.
	address := fmt.Sprintf("%s:%d", "0.0.0.0", 5003)
	network := "tcp"

	listenNet, err := net.Listen(network, address)
	if err != nil {
		slog.Error("Failed to listen to address", err, "network", network, "address", address)
		cancel()
		<-ctx.Done()
	}

	slog.Info("🌏 start server...", "address", address)

	defer func() {
		if err := listenNet.Close(); err != nil {
			slog.Error("failed to close", err, "network", network, "address", address)
			<-ctx.Done()
		}
	}()

	err = server.Serve(listenNet)
	if err != nil {
		slog.Error("Failed start gRPC server", err, "network", network, "address", address)
		cancel()
		<-ctx.Done()
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGILL)

	select {
	case v := <-quit:
		slog.Info("signal.Notify", v)
	case done := <-ctx.Done():
		slog.Info("ctx.Done", "app done", done)
	}

	return nil
}
