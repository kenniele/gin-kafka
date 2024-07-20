package main

import (
	"github.com/kenniele/messagio-gin-postgresql-kafka/backend/config"
	"log"
)

func main() {
	if cfg, err := config.LoadConfig("config/config.yaml"); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
}
