package models

import (
	"messagio-gin-postrgresql-kafka/internal/db"
)

type Message struct {
	ID   int    `json:"id"`
	Data string `json:"data"`
}

func (message Message) Create() (int, error) {
	id, err := db.CreateMessage(message.Data)
	if err != nil {
		return 0, err
	}
	message.ID = id
	return id, nil
}
