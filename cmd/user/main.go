package main

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"

	httpServer "github.com/duongnln96/blog-realworld/internal/user/app/http_server"
	migrate "github.com/duongnln96/blog-realworld/internal/user/app/migrations"
	genOpenApi3 "github.com/duongnln96/blog-realworld/internal/user/app/open_api3"

	"github.com/duongnln96/blog-realworld/pkg/config"

	"github.com/urfave/cli/v2"
)

func main() {

	// set max process
	runtime.GOMAXPROCS(runtime.NumCPU())

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
		{
			Name:        "open-api",
			Description: "Generate OpenAPI3 doc",
			Action: func(c *cli.Context) error {
				return genOpenApi3.Run(c)
			},
			Flags: []cli.Flag{
				&cli.PathFlag{
					Name:     "output",
					Usage:    "output path",
					Required: true,
				},
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
