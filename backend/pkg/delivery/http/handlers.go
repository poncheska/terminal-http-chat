package http

import (
	"encoding/json"
	"github.com/poncheska/terminal-http-chat/backend/pkg/utils"
	"net/http"
)

var chatNotifiers map[string]utils.ChatNotifier

type ChatUser struct {
	Token string `json:"token"`
}

type SignInData struct {
	Username string `json:"login"`
	Password string `json:"password"`
}

func (s http2.Server) ChatHandler(w http.ResponseWriter, r *http.Request) {
	var user ChatUser
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		http2.WriteErrorResponse(w, err)
		return
	}
}

func (s http2.Server) ChatsHandler(w http.ResponseWriter, r *http.Request) {
	var user ChatUser
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		http2.WriteErrorResponse(w, err)
		return
	}
}

func (s http2.Server) SignInHandler(w http.ResponseWriter, r *http.Request) {
	data := &SignInData{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		http2.WriteErrorResponse(w, err)
		return
	}
	user, err := s.store.User.GetByCredentials(data.Username, data.Password)
	if err != nil {
		http2.WriteErrorResponse(w, err)
		return
	}

}

func (s http2.Server) SignUpHandler(w http.ResponseWriter, r *http.Request) {

}
