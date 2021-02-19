package store

import (
	"fmt"
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
	var chats []models.Chat
	err := cs.db.Select(&chats, "SELECT * FROM chat")
	if err != nil {
		return []models.Chat{}, err
	}
	return chats, nil
}

func (cs *ChatStore) Create(chat models.Chat) (int64, error) {
	var id int64
	err := cs.db.QueryRow("INSERT INTO chat(name, admin_id) VALUES ($1,$2) RETURNING id",
		chat.Name, chat.AdminId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (cs *ChatStore) Delete(chatId, adminId int64) error {
	res, err := cs.db.Exec("DELETE FROM chat WHERE id = $1 AND admin_id = $2", chatId, adminId)
	if err != nil {
		return err
	}
	if ra, err := res.RowsAffected(); err != nil {
		return err
	}else if ra == 0{
		return fmt.Errorf("chat have not deleted, may be you are not admin")
	}
	return nil
}
