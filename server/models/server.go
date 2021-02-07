package models

import (
	"github.com/jmoiron/sqlx"
)

type Server struct {
	DB    sqlx.DB
	Chats []Chat
}

type Chat struct {
	Id int64
	Name string
	Messages []Message
}

type Message struct{
	Sender User
	Text   string
}

type User struct{
	Id int64
	Nickname string
}
