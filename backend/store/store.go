package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/poncheska/terminal-http-chat/backend/models"
)

type Store struct {
	Chat
	Message
	User
}

type Message interface {
	GetAll(chatId int64) ([]models.Message, error)
	Create(message models.Message) error
}

type Chat interface {
	GetAll() ([]models.Chat, error)
	Create(chat models.Chat) error
}

type User interface {
	GetById(userId int64) (models.User, error)
	Create(login, password string) error
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		Chat: NewChatStore(db),
		Message: NewMessageStore(db),
		User: NewUserStore(db),
	}
}
