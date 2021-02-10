package store

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type Store struct {
	DB *sqlx.DB
}

func NewStore(connStr string) Store{
	db,err := sqlx.Connect("postgres",connStr)
	if err != nil{
		log.Fatal(err.Error())
	}
	return Store{db}
}

func (s Store) CheckUserData(login, password string) bool{
	return false
}

func (s Store) CreateAccount(login, password string) error{
	return fmt.Errorf("error")
}
