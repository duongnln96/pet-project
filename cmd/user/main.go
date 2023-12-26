package main

import (
	"fmt"
	"log/slog"
	"os"

	httpServer "github.com/duongnln96/blog-realworld/internal/user/app/http_server"
	migrate "github.com/duongnln96/blog-realworld/internal/user/app/migrations"
	"github.com/duongnln96/blog-realworld/pkg/config"

	"github.com/urfave/cli/v2"
)

func main() {
	configPath, _ := os.Getwd()
	configs := config.LoadConfig(fmt.Sprintf("%s/config/user", configPath))

	app := cli.NewApp()

	app.Commands = []*cli.Command{
		{
			Name:    "api",
			Aliases: []string{"s"},
			Action: func(c *cli.Context) error {
				return httpServer.Run(configs)
			},
		},
		{
			Name: "migration",
			Action: func(c *cli.Context) error {
				return migrate.RunMigrations(configs)
			},
		},
	}

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "s")
	}

	if err := app.Run(os.Args); err != nil {
		slog.Error("Server application running error", "err_info", err.Error())
		return
	}
}
