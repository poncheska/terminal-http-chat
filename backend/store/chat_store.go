package store

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/poncheska/terminal-http-chat/backend/models"
)

type ChatStore struct {
	db *sqlx.DB
}

func NewChatStore(db *sqlx.DB) *ChatStore{
	return &ChatStore{db}
}

func (cs *ChatStore) GetAll() ([]models.Chat, error){
	return []models.Chat{}, fmt.Errorf("")
}

func (cs *ChatStore) Create(chat models.Chat) error{
	return fmt.Errorf("")
}