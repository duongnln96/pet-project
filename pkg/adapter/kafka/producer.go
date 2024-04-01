package kafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaProducer struct {
	producer *kafka.Producer
}

func NewKafkaProducer(cfg KafkaProducerConfig) (*KafkaProducer, error) {

	producer, err := kafka.NewProducer(&cfg.Configs)
	if err != nil {
		return nil, fmt.Errorf("kafka.NewProducer error %s", err.Error())
	}

	// Phải đọc events ra liên tục, nếu ko sẽ bị memory leak
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition.Error)
				}
			default:
				fmt.Printf("Ignored event: %s\n", ev)
			}
		}
	}()

	return &KafkaProducer{
		producer,
	}, nil
}
