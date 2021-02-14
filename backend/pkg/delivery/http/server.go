package http

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"github.com/poncheska/terminal-http-chat/backend/pkg/store"
	"github.com/poncheska/terminal-http-chat/backend/pkg/utils"
	"net/http"
)

type Handler struct {
	store        *store.Store
	chatNotifier *utils.ChatNotifier
	tokenService *utils.TokenService
}

type JSONError struct {
	Msg string `json:"message"`
}

func NewServer(db *sqlx.DB, jwtKey string) Handler {
	return Handler{
		store:        store.NewStore(db),
		chatNotifier: utils.NewChatNotifier(),
		tokenService: utils.NewTokenService(jwtKey),
	}
}

func NewJSONError(msg string) JSONError {
	return JSONError{msg}
}

func WriteErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(JSONError{err.Error()})
}

func WriteUnauthorizedResponse(w http.ResponseWriter, msg string){
	w.WriteHeader(http.StatusUnauthorized)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(JSONError{msg})
}
