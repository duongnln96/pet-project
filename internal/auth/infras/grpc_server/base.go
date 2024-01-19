package grpc_server

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/duongnln96/blog-realworld/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcServerAdapterI interface {
	GetServerInstance() *grpc.Server
	Start() error
	Stop()
}

type grpcServerAdapter struct {
	configs *config.Configs
	logger  *slog.Logger

	s *grpc.Server
}

func NewGRPCServerAdapter(
	configs *config.Configs,
	logger *slog.Logger,

	serverOptions ...grpc.ServerOption,
) GrpcServerAdapterI {
	server := grpc.NewServer(
		serverOptions...,
	)

	reflection.Register(server)

	return &grpcServerAdapter{
		configs: configs,
		logger:  logger,
		s:       server,
	}
}

func (a *grpcServerAdapter) GetServerInstance() *grpc.Server {
	return a.s
}

func (a *grpcServerAdapter) Start() error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		defer a.s.GracefulStop()
		<-ctx.Done()
	}()

	// gRPC Server.
	address := fmt.Sprintf("%s:%d", a.configs.Other.Get("grpc_address"), a.configs.Other.Get("grpc_port"))
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

	err = a.s.Serve(listenNet)
	if err != nil {
		slog.Error("Failed start gRPC server", err, "network", network, "address", address)
		cancel()
		<-ctx.Done()
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGILL)

	select {
	case v := <-quit:
		slog.Info("Receieved signal", v)
	case done := <-ctx.Done():
		slog.Info("ctx.Done", "app done", done)
	}

	return nil
}

func (a *grpcServerAdapter) Stop() {
	a.s.GracefulStop()
}
