package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/poncheska/terminal-http-chat/backend/pkg/models"
)

type ChatStore struct {
	db *sqlx.DB
}

func NewChatStore(db *sqlx.DB) *ChatStore {
	return &ChatStore{db}
}

func (cs *ChatStore) GetAll() ([]models.Chat, error) {
	var chat []models.Chat
	err := cs.db.Get(&chat, "SELECT * FROM chat")
	if err != nil {
		return []models.Chat{}, err
	}
	return chat, nil
}

func (cs *ChatStore) Create(chat models.Chat) (int64, error) {
	res, err := cs.db.Exec("INSERT INTO chat(name) VALUES ($1,$2) RETURNING id", chat.Name)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
