package postgres

import (
	"encoding/json"
	"log"
	"time"
)

type ScyllaDBConfig struct {
	Name      string        `mapstructure:"name" json:"name"`
	Host      string        `mapstructure:"host" json:"host"`
	Port      int           `mapstructure:"port" json:"port"`
	UserName  string        `mapstructure:"username" json:"username"`
	Password  string        `mapstructure:"password" json:"password"`
	DBName    string        `mapstructure:"dbname" json:"dbname"`
	Timeout   time.Duration `mapstructure:"timeout" json:"timeout"`
	PoolLimit int           `mapstructure:"pool_limit" json:"pool_limit"`
}

func (m ScyllaDBConfig) PrettyPrint() string {
	bytes, _ := json.MarshalIndent(m, "", " ")
	return string(bytes)
}

type ScyllaDBConfigMap map[string]*ScyllaDBConfig

func (c ScyllaDBConfigMap) Get(name string) *ScyllaDBConfig {
	config, ok := c[name]
	if !ok {
		log.Panicf("Not found config name: [%s]", name)
	}

	return config
}
