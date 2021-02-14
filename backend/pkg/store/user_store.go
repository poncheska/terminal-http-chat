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

func (us *UserStore) GetByCredentials(login, password string) (models.User, error){
	return models.User{},fmt.Errorf("")
}

func (us *UserStore) Create(login, password string) error{
	return fmt.Errorf("")
}