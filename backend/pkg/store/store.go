package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/poncheska/terminal-http-chat/backend/pkg/models"
)

type Store struct {
	Chat
	Message
	User
	View
}

type Message interface {
	Create(message models.Message) (int64, error)
}

type Chat interface {
	GetAll() ([]models.Chat, error)
	Create(chat models.Chat) (int64, error)
	Delete(chatId, adminId int64) error
}

type User interface {
	GetByCredentials(username, password string) (models.User, error)
	Create(username, password string) (int64, error)
}

type View interface {
	GetAllMessageData(chatId int64) ([]models.MessageData, error)
	GetMessageDataById(msgId int64) (models.MessageData, error)
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		Chat:    NewChatStore(db),
		Message: NewMessageStore(db),
		User:    NewUserStore(db),
		View:    NewViewStore(db),
	}
}
