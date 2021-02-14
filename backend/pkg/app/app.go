package app

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	http2 "github.com/poncheska/terminal-http-chat/backend/pkg/delivery/http"
	"log"
	"net/http"
	"os"
)

var jwtKey = "gdfgdfg"

func Run(){
	db, err := sqlx.Connect("postgres","user=admin password=password sslmode=disable")
	if err != nil{
		log.Fatal(err)
	}
	srv := http2.NewServer(db, jwtKey)

	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}

	r := mux.NewRouter()
	r.HandleFunc("/signin", srv.SignIn)
	r.HandleFunc("/signup", srv.SignUp)
	r.HandleFunc("/chats", srv.AuthChecker(srv.Chats))
	r.HandleFunc("/chat/{id:[0-9]+}", srv.AuthChecker(srv.Chat))
	log.Println("Handler started")
	log.Fatal(http.ListenAndServe(":"+port, r))
}