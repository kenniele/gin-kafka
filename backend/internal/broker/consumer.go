package broker

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"messagio-gin-postrgresql-kafka/config"
)

func NewConsumer(cfg config.KafkaConfig) *kafka.Consumer {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "87.228.13.94:9092",
		"group.id":          cfg.GroupID,
		"auto.offset.reset": "smallest",
	})

	if err != nil {
		log.Fatalf("Error - failed to create consumer: %s\n", err)
	}

	err = consumer.Subscribe(cfg.Topic, nil)
	if err != nil {
		log.Fatalf("Error - failed to subscribe to topic: %s\n", err)
	}

	return consumer
}

func ProcessMessage(consumer *kafka.Consumer) {
	for {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			log.Printf("Message on %s: %s\n", e.TopicPartition, string(e.Value))
		case kafka.Error:
			log.Printf("Error on %s: %s\n", e.Code(), e.Error())
		}
	}
}
