package main

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"

	"github.com/duongnln96/blog-realworld/pkg/config"
	"github.com/duongnln96/blog-realworld/pkg/logger"
	"github.com/sirupsen/logrus"

	grpcServer "github.com/duongnln96/blog-realworld/internal/auth/adapter/grpc_server"
)

func main() {

	// set max process
	runtime.GOMAXPROCS(runtime.NumCPU())

	configPath, _ := os.Getwd()
	configs := config.LoadConfig(fmt.Sprintf("%s/config/auth", configPath))

	// setup logger
	logrus.SetLevel(logger.ConvertLogLevel("debug"))
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)

	logger := slog.New(logger.NewLogrusHandler(logrus.StandardLogger()))

	app, cancel := grpcServer.InitNewApp(configs, logger)
	defer cancel()

	if err := app.Serve(); err != nil {
		slog.Error("Application running error", "err_info", err.Error())
		return
	}
}
