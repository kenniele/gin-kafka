package main

import (
	"github.com/joho/godotenv"
	"log"
	"messagio-gin-postrgresql-kafka/config"
	"messagio-gin-postrgresql-kafka/internal/broker"
	"messagio-gin-postrgresql-kafka/internal/db"
	"messagio-gin-postrgresql-kafka/internal/handlers"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("Error loading .env file")
		return
	}
}

func main() {
	cfg := config.NewConfig()

	op := broker.NewOrderPlacer(cfg.Kafka)

	_, err := db.Connect(cfg.DB)
	db.InitKafka(op)

	if err != nil {
		log.Fatalf("Error - connecting to Database: %v", err)
		return
	}
	defer db.Close()

	consumer := broker.NewConsumer(cfg.Kafka)

	go broker.ProcessMessage(consumer)

	router := handlers.CreateRouter()
	if router == nil {
		log.Fatalf("Error - creating router: %v", err)
	}

	handlers.Run(router)
}
