package main

import (
	"fmt"
	"os"

	"github.com/duongnln96/blog-realworld/pkg/config"
)

func main() {
	configPath, _ := os.Getwd()
	configs := config.LoadConfig(fmt.Sprintf("%s/config/user", configPath))

	fmt.Println(configs.ScyllaDBConfigMap.Get("scylladb_standalone"))
}
