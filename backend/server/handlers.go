package server

import (
	"encoding/json"
	"github.com/poncheska/terminal-http-chat/backend/utils"
	"net/http"
)

var chatNotifiers map[string]utils.ChatNotifier

type ChatUser struct{
	Token string `json:"token"`
}

type SignInData struct{
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (s Server) ChatHandler(w http.ResponseWriter, r *http.Request){
	var user ChatUser
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (s Server) ChatsHandler(w http.ResponseWriter, r *http.Request){
	var user ChatUser
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (s Server) SignInHandler(w http.ResponseWriter, r *http.Request){
	data := &SignInData{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (s Server) SignUpHandler(w http.ResponseWriter, r *http.Request){

}
