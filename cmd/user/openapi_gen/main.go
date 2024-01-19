package main

import (
	"log/slog"
	"os"
	"runtime"

	genOpenApi3 "github.com/duongnln96/blog-realworld/internal/user/adapter/openapi3_gen"

	"github.com/urfave/cli/v2"
)

func main() {

	// set max process
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := cli.NewApp()

	app.Commands = []*cli.Command{
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
