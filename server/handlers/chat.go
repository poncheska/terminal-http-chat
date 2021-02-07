package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"terminal-http-chat/server/utils"
)

var chatNotifiers map[string]utils.ChatNotifier

type ChatUser struct{
	Token string `json:"token"`
}

func ChatHandler(w http.ResponseWriter, r *http.Request){
	var user ChatUser
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		log.Fatal(err.Error())
	}
}

