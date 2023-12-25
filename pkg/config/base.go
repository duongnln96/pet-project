package config

import (
	"log"
	"sync"

	scyllaDB "github.com/duongnln96/blog-realworld/pkg/adapter/scylladb"
	"github.com/spf13/viper"
)

var (
	onceLoadConfig = sync.Once{}
)

var onceConfigs *Configs

type Configs struct {
	ScyllaDBConfigMap scyllaDB.ScyllaDBConfigMap
}

func LoadConfig(serviceConfigPath string) *Configs {
	log.Println("Load config from path ", serviceConfigPath)

	var onceConfigs = new(Configs)

	readConfig(serviceConfigPath, "scylladb", &(onceConfigs.ScyllaDBConfigMap))

	return onceConfigs
}

func readConfig(configPath, configName string, result interface{}) error {

	viperI := viper.New()

	viperI.AutomaticEnv()
	viperI.AddConfigPath(configPath)
	viperI.SetConfigType("yaml")
	viperI.SetConfigName(configName)

	err := viperI.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Printf("[ERROR] %s.yaml file not found", configName)
			err = nil
		} else {
			// Config file was found but another error was produced
			log.Panicf("[ERROR] reading %s.yaml %s", configName, err.Error())
		}
	}

	err = viperI.Unmarshal(&result)
	if err != nil {
		log.Printf("viperI.Unmarshal %s %s", configName, err.Error())
		return err
	}

	return nil
}
