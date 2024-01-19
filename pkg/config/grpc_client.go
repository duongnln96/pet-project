package config

import (
	"encoding/json"
	"log"
	"time"
)

type GrpcClientConfig struct {
	Url     string        `mapstructure:"url" json:"url"`
	Timeout time.Duration `mapstructure:"timeout" json:"timeout"`
}

func (m GrpcClientConfig) PrettyPrint() string {
	bytes, _ := json.MarshalIndent(m, "", " ")
	return string(bytes)
}

type GrpcClientConfigMap map[string]*GrpcClientConfig

func (c GrpcClientConfigMap) Get(name string) *GrpcClientConfig {
	config, ok := c[name]
	if !ok {
		log.Panicf("Not found config name: [%s]", name)
	}

	return config
}
