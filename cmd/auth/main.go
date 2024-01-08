package main

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"

	"github.com/duongnln96/blog-realworld/pkg/config"

	grpcServer "github.com/duongnln96/blog-realworld/internal/auth/app/grpc_server"
	migrate "github.com/duongnln96/blog-realworld/internal/auth/app/migrations"

	"github.com/urfave/cli/v2"
)

func main() {

	// set max process
	runtime.GOMAXPROCS(runtime.NumCPU())

	configPath, _ := os.Getwd()
	configs := config.LoadConfig(fmt.Sprintf("%s/config/auth", configPath))

	app := cli.NewApp()

	app.Commands = []*cli.Command{
		{
			Name:    "grpc_server",
			Aliases: []string{"grpc"},
			Action: func(c *cli.Context) error {
				return grpcServer.RunGRPCServer(configs)
			},
		},
		{
			Name: "migrations",
			Action: func(c *cli.Context) error {
				return migrate.Run(configs)
			},
		},
	}

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "grpc")
	}

	if err := app.Run(os.Args); err != nil {
		slog.Error("Server application running error", "err_info", err.Error())
		return
	}
}
