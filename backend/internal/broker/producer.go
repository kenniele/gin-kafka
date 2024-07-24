package broker

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"messagio-gin-postrgresql-kafka/config"
)

type OrderPlacer struct {
	producer     *kafka.Producer
	topic        string
	deliveryChan chan kafka.Event
}

func NewOrderPlacer(cfg config.KafkaConfig) *OrderPlacer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "87.228.13.94:9092",
		"client.id":         cfg.GroupID,
		"acks":              "all",
	})

	if err != nil {
		log.Printf("Failed to create producer: %s\n", err)
	}
	return &OrderPlacer{
		producer:     p,
		topic:        cfg.Topic,
		deliveryChan: make(chan kafka.Event, 10000),
	}
}

func (op *OrderPlacer) PlaceOrder(orderType string) error {
	var (
		payload = []byte(orderType)
	)
	err := op.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &op.topic,
			Partition: kafka.PartitionAny,
		},
		Value: payload,
	},
		op.deliveryChan,
	)

	if err != nil {
		return err
	}
	<-op.deliveryChan
	return nil
}
