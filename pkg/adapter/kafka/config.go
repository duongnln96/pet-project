package kafka

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaConfig struct {
	BootstrapServers []string                       `mapstructure:"bootstrap_servers" yaml:"bootstrap_servers"`
	Configs          map[string]interface{}         `mapstructure:"configs" yaml:"configs"`
	Producers        map[string]KafkaProducerConfig `mapstructure:"producers" yaml:"producers"`
	Consumers        map[string]KafkaConsumerConfig `mapstructure:"consumers" yaml:"consumers"`
}

func (m KafkaConfig) PrettyPrint() string {
	bytes, _ := json.MarshalIndent(m, "", " ")
	return string(bytes)
}

func (m KafkaConfig) GetProducerConfig(name string) KafkaProducerConfig {
	var configMap = make(kafka.ConfigMap)

	// lấy config theo name producer
	producer, ok := m.Producers[name]
	if !ok {
		log.Panicf("Not found kafka producer config name %s", name)
	}

	// lấy config chung
	if len(m.Configs) > 0 {
		for k, v := range m.Configs {
			configMap[k] = v
		}
	}

	// lấy config từ producer map vào
	for k, v := range producer.Configs {
		configMap[k] = v
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Panicf("hostname: [%s]", name)
	}

	// overide một số config riêng
	configMap["bootstrap.servers"] = strings.Join(m.BootstrapServers, ",")
	configMap["client.id"] = hostname

	// set thêm giá trị mặc định nếu chưa có
	if _, ok := configMap["compression.type"]; !ok {
		configMap["compression.type"] = "gzip"
	}
	if _, ok := configMap["acks"]; !ok {
		configMap["acks"] = "1"
	}
	if _, ok := configMap["linger.ms"]; !ok {
		configMap["linger.ms"] = 200
	}

	return KafkaProducerConfig{
		Configs: configMap,
	}
}

func (m KafkaConfig) GetConsumerConfig() KafkaConsumerConfig {
	return KafkaConsumerConfig{}
}

type KafkaConfigs map[string]KafkaConfig

func (c KafkaConfigs) Get(name string) KafkaConfig {
	config, ok := c[name]
	if !ok {
		log.Panicf("Not found kafka config name %s", name)
	}

	return config
}

type KafkaProducerConfig struct {
	Configs kafka.ConfigMap `mapstructure:"configs" yaml:"configs"`
}

func (m KafkaProducerConfig) PrettyPrint() string {
	bytes, _ := json.MarshalIndent(m, "", " ")
	return string(bytes)
}

type KafkaConsumerConfig struct {
	Configs kafka.ConfigMap `mapstructure:"configs" yaml:"configs"`
}

func (m KafkaConsumerConfig) PrettyPrint() string {
	bytes, _ := json.MarshalIndent(m, "", " ")
	return string(bytes)
}

type KafkaAdminConfig struct {
	Configs kafka.ConfigMap `mapstructure:"configs" yaml:"configs"`
}

func (m KafkaAdminConfig) PrettyPrint() string {
	bytes, _ := json.MarshalIndent(m, "", " ")
	return string(bytes)
}
