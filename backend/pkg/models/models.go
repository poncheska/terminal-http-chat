package models

import "time"

type Chat struct {
	Id   int64  `json:"id" ,db:"id"`
	Name string `json:"name" ,db:"name"`
}

type Message struct {
	Id       int64     `json:"id" ,db:"id"`
	SenderId int64     `json:"user_id" ,db:"user_id"`
	Date     time.Time `json:"date" ,db:"date"`
	Text     string    `json:"text" ,db:"text"`
}

type MessageData struct {
	SenderName string    `json:"name" ,db:"name"`
	Date       time.Time `json:"date" ,db:"date"`
	Text       string    `json:"text" ,db:"text"`
}

type User struct {
	Id       int64  `json:"id" ,db:"id"`
	Nickname string `json:"name" ,db:"name"`
	Password string `json:"password" ,db:"password"`
}
