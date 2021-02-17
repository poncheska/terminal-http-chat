package store

import (
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
	var md []models.MessageData
	err := vs.db.Get(&md, "SELECT name, date, text FROM message_data WHERE chat_id = $1", chatId)
	if err != nil {
		return []models.MessageData{}, err
	}
	return md, nil
}

func (vs *ViewStore) GetMessageDataById(msgId int64) (models.MessageData, error) {
	var md models.MessageData
	err := vs.db.Get(&md, "SELECT name, date, text FROM message_data WHERE id = $1", msgId)
	if err != nil {
		return models.MessageData{}, err
	}
	return md, nil
}
