package main

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"

	"github.com/duongnln96/blog-realworld/pkg/config"

	migrate "github.com/duongnln96/blog-realworld/internal/auth/adapter/migrations"
)

func main() {

	// set max process
	runtime.GOMAXPROCS(runtime.NumCPU())

	configPath, _ := os.Getwd()
	configs := config.LoadConfig(fmt.Sprintf("%sconfig/auth", configPath))

	if err := migrate.Run(configs); err != nil {
		slog.Error("Server application running error", "err_info", err.Error())
		return
	}
}
