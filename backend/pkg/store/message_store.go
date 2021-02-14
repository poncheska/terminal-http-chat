package store

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/poncheska/terminal-http-chat/backend/pkg/models"
)

type MessageStore struct {
	db *sqlx.DB
}

func NewMessageStore(db *sqlx.DB) *MessageStore {
	return &MessageStore{db}
}

func (ms *MessageStore) Create(message models.Message) (int64, error) {
	return 0, fmt.Errorf("")
}
