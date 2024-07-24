package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"messagio-gin-postrgresql-kafka/config"
	"messagio-gin-postrgresql-kafka/internal/broker"
)

var placer *broker.OrderPlacer
var DataBase *sql.DB

func InitKafka(op *broker.OrderPlacer) {
	placer = op
}

func Connect(cfg config.DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database)
	DataBase, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to PostgreSQL server")

	return DataBase, nil
}

func CreateMessage(Data string) (int, error) {
	if DataBase == nil {
		cfg := config.NewConfig()
		DB, err := Connect(cfg.DB)
		if err != nil {
			log.Printf("Error - connecting to database: %s", err)
		}
		DataBase = DB
	}
	row := DataBase.QueryRow(`INSERT INTO message (data) VALUES ($1) RETURNING "id"`, Data)
	log.Printf("Message data - %s\n", Data)
	var ID int
	if err := row.Scan(&ID); err != nil {
		log.Printf("Error - creating message: %v", err)
		return 0, errors.New("ошибка во время создания сообщения")
	}
	log.Printf("Message created with ID: %d", ID)
	err := placer.PlaceOrder(Data)
	if err != nil {
		return 0, err
	}

	return ID, nil
}

func GetAllMessages() ([]string, error) {
	rows, err := DataBase.Query(`SELECT data FROM message`)
	if err != nil {
		log.Printf("Error - getting all messages: %v", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		if err = rows.Close(); err != nil {
			log.Printf("Error - closing rows: %v", err)
		}
	}(rows)

	var messages []string
	for rows.Next() {
		var message string
		if err = rows.Scan(&message); err != nil {
			log.Printf("Error - getting all messages: %v", err)
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func Close() {
	if err := DataBase.Close(); err != nil {
		log.Fatalf("Error - closing DB connection: %v", err)
	}
}
