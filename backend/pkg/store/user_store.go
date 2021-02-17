package store

import (
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
	var user models.User
	err := us.db.Get(&user, "SELECT * FROM users WHERE name = $1 AND password = $2", username, password)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (us *UserStore) Create(username, password string) (int64, error) {
	res, err := us.db.Exec("INSERT INTO users(name,password) VALUES ($1,$2) RETURNING id", username, password)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
