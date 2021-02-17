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
	res, err := cs.db.Exec("INSERT INTO chat(name, admin_id) VALUES ($1,$2) RETURNING id", chat.Name, chat.AdminId)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (cs *ChatStore) Delete(chatId, adminId int64) error {
	_, err := cs.db.Exec("DELETE FROM chat WHERE id = $1 AND admin_id = $2", chatId, adminId)
	if err != nil {
		return err
	}
	return nil
}
