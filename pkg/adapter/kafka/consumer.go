package kafka

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

type KafkaConsumer struct {
	consumer *kafka.Consumer
}

func NewKafkaConsumer(cfg KafkaConsumerConfig) *KafkaConsumer {

	return &KafkaConsumer{
		consumer: &kafka.Consumer{},
	}
}
