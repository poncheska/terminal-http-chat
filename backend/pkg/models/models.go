package models

import "time"

type Chat struct {
	Id   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	AdminId int64 `json:"-" db:"admin_id"`
}

type Message struct {
	Id       int64     `json:"id" db:"id"`
	SenderId int64     `json:"user_id" db:"user_id"`
	ChatId   int64     `json:"chat_id" db:"chat_id"`
	Date     time.Time `json:"date" db:"date"`
	Text     string    `json:"text" db:"text"`
}

type MessageData struct {
	SenderName string    `json:"name" db:"name"`
	Date       time.Time `json:"date" db:"date"`
	Text       string    `json:"text" db:"text"`
}

type User struct {
	Id       int64  `json:"id" db:"id"`
	Username string `json:"name" db:"name"`
	Password string `json:"password" db:"password"`
}
