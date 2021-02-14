package models

type Chat struct {
	Id       int64
	Name     string
	Messages []Message
}

type Message struct {
	Sender User
	Text   string
}

type User struct {
	Id       int64
	Nickname string
	Password string
}
