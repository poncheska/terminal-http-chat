package store

import (
"fmt"
"github.com/jmoiron/sqlx"
"github.com/poncheska/terminal-http-chat/backend/models"
)

type MessageStore struct {
	db *sqlx.DB
}

func NewMessageStore(db *sqlx.DB) *MessageStore{
	return &MessageStore{db}
}

func (ms *MessageStore) GetAll(chatId int64) ([]models.Message, error){
	return []models.Message{}, fmt.Errorf("")
}

func (ms *MessageStore) Create(message models.Message) error{
	return fmt.Errorf("")
}
