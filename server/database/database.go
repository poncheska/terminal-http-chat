package database

import (
	"github.com/jmoiron/sqlx"
	"log"
)

var db *sqlx.DB

func InitDB(connStr string){
	var err error
	db,err = sqlx.Connect("postgres",connStr)
	if err != nil{
		log.Fatal(err.Error())
	}
}
