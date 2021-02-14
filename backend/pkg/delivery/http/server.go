package http

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"github.com/poncheska/terminal-http-chat/backend/pkg/store"
	"github.com/poncheska/terminal-http-chat/backend/pkg/utils"
	"net/http"
)

type Server struct {
	store         *store.Store
	chatNotifiers map[string]utils.ChatNotifier
	tokenService  *utils.TokenService
}

type JSONError struct {
	Msg string `json:"error"`
}

func NewServer(db *sqlx.DB, jwtKey string) Server {
	return Server{
		store:         store.NewStore(db),
		chatNotifiers: map[string]utils.ChatNotifier{},
		tokenService:  utils.NewTokenService(jwtKey),
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
