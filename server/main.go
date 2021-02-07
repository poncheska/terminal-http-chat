package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"terminal-http-chat/server/database"
	"terminal-http-chat/server/handlers"
)

func main() {
	database.InitDB("user=admin password=password sslmode=disable")

	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}

	r := mux.NewRouter()
	r.HandleFunc("/signin", handlers.SignInHandler)
	r.HandleFunc("/signup", handlers.SignUpHandler)
	r.HandleFunc("/chat", handlers.ChatHandler)
	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
