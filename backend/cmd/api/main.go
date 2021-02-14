package main

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/poncheska/terminal-http-chat/backend/server"
	"log"
	"net/http"
	"os"
)

func main() {
	db, err := sqlx.Connect("postgres","user=admin password=password sslmode=disable")
	if err != nil{
		log.Fatal(err)
	}
	srv := server.NewServer(db)

	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}

	r := mux.NewRouter()
	r.HandleFunc("/signin", srv.SignInHandler)
	r.HandleFunc("/signup", srv.SignUpHandler)
	r.HandleFunc("/chats", server.AuthChecker(srv.ChatsHandler))
	r.HandleFunc("/chat/$id", server.AuthChecker(srv.ChatHandler))
	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
