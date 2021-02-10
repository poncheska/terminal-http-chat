package server

import (
	"github.com/jmoiron/sqlx"
	"github.com/poncheska/terminal-http-chat/backend/store"
	"github.com/poncheska/terminal-http-chat/backend/utils"
)

type Server struct {
	store         *store.Store
	chatNotifiers map[string]utils.ChatNotifier
}

func NewServer(db *sqlx.DB) Server {
	return Server{
		store: store.NewStore(db),
		chatNotifiers: map[string]utils.ChatNotifier{},
	}
}
