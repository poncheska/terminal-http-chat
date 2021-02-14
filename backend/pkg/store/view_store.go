package store

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/poncheska/terminal-http-chat/backend/pkg/models"
)

type ViewStore struct {
	db *sqlx.DB
}

func NewViewStore(db *sqlx.DB) *ViewStore {
	return &ViewStore{db}
}

func (vs *ViewStore) GetAllMessageData(chatId int64) ([]models.MessageData, error) {
	return []models.MessageData{}, fmt.Errorf("")
}

func (vs *ViewStore) GetMessageDataById(msgId int64) (models.MessageData, error) {
	return models.MessageData{}, fmt.Errorf("")
}
