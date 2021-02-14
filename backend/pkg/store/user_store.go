package store

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/poncheska/terminal-http-chat/backend/pkg/models"
)

type UserStore struct {
	db *sqlx.DB
}

func NewUserStore(db *sqlx.DB) *UserStore {
	return &UserStore{db}
}

func (us *UserStore) GetByCredentials(username, password string) (models.User, error) {
	return models.User{}, fmt.Errorf("")
}

func (us *UserStore) Create(username, password string) (int64, error) {
	return 0, fmt.Errorf("")
}
