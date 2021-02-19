package http

import "fmt"

type AuthData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateChatReq struct {
	ChatName string `json:"chat_name"`
}

func (ad AuthData) CheckNotEmpty() error{
	if ad.Password == "" || ad.Username == ""{
		return fmt.Errorf("username or password is empty")
	}
	return nil
}

func (c CreateChatReq) CheckNotEmpty() error{
	if c.ChatName == "" {
		return fmt.Errorf("chat_name is empty")
	}
	return nil
}