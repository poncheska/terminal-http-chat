package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	http2 "github.com/poncheska/terminal-http-chat/backend/pkg/delivery/http"
	"log"
	"net/http"
	"os"
)

var jwtKey = "gdfgdfg"

func Run(){
	host := os.Getenv("POSTGRES_HOST")
	if host == ""{
		host = "127.0.0.1"
	}

	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v?sslmode=%v",
		"postgres",
		"password",
		host,
		"5432",
		"disable")
	db, err := sqlx.Connect("postgres",connStr)
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
	r.HandleFunc("/chat/{id:[0-9]+}/delete", srv.AuthChecker(srv.DeleteChat))
	r.HandleFunc("/chat/create", srv.AuthChecker(srv.CreateChat))
	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":"+port, r))
}