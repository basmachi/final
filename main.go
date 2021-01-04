package main

import (
	"database/sql"
	"final/cmd/app"
	"final/pkg/core"
	"fmt"
	"github.com/julienschmidt/httprouter"
	_"github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	DB, err :=sql.Open("sqlite3","myfinal")
	if err != nil {
		log.Fatal("can't find sql connection", err)
	}
	fmt.Println("Connection to db is okay")
	svc := core.NewUserSvc(DB)
	router := httprouter.New()
	server := app.NewMainServer(DB, router, svc)

	server.Start(DB)
}
