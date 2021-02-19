package store

import (
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
	var id int64
	err := ms.db.QueryRow(
		"INSERT INTO message(user_id, chat_id, date, text) VALUES ($1,$2,$3,$4) RETURNING id",
		message.SenderId, message.ChatId, message.Date, message.Text).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
