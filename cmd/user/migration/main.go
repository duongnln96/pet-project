package main

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"

	migrate "github.com/duongnln96/blog-realworld/internal/user/adapter/migrations"

	"github.com/duongnln96/blog-realworld/pkg/config"
)

func main() {
	// set max process
	runtime.GOMAXPROCS(runtime.NumCPU())

	configPath, _ := os.Getwd()
	configs := config.LoadConfig(fmt.Sprintf("%s/config/user", configPath))

	if err := migrate.RunMigrations(configs); err != nil {
		slog.Error("Application running error", "err_info", err.Error())
		return
	}
}
