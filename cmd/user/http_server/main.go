package main

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"

	httpServer "github.com/duongnln96/blog-realworld/internal/user/adapter/http_server"

	"github.com/duongnln96/blog-realworld/pkg/config"
)

func main() {

	// set max process
	runtime.GOMAXPROCS(runtime.NumCPU())

	configPath, _ := os.Getwd()
	configs := config.LoadConfig(fmt.Sprintf("%s/config/user", configPath))

	app, cancel, err := httpServer.InitNewApp(configs)
	defer cancel()
	if err != nil {
		slog.Error("Application initialize error", "err_info", err.Error())
		return
	}

	if err := app.Serve(); err != nil {
		slog.Error("Application running error", "err_info", err.Error())
		return
	}
}
