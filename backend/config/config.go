package config

import "os"

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type KafkaConfig struct {
	Brokers []string
	Topic   string
	GroupID string
}

type Config struct {
	DB    DBConfig
	Kafka KafkaConfig
}

func NewConfig() *Config {
	return &Config{
		DB: DBConfig{
			Host:     getEnv("HOST", "localhost"),
			Port:     getEnv("PORT", "5432"),
			Username: getEnv("USER", "postgres"),
			Password: getEnv("PASSWORD", "postgres"),
			Database: getEnv("DATABASE", "postgres"),
		},
		Kafka: KafkaConfig{
			Brokers: []string{"localhost:9092"},
			Topic:   getEnv("TOPIC", "test"),
			GroupID: getEnv("GROUP_ID", "test"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
