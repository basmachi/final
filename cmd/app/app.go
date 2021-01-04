package app

import (
	"database/sql"
	"final/db"
	"final/pkg/core"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type MainServer struct {
	db *sql.DB
	router  *httprouter.Router
	svc *core.UserSvc
}

func NewMainServer(db *sql.DB, router *httprouter.Router, svc *core.UserSvc) *MainServer {
	return &MainServer{db: db, router: router, svc: svc}
}


func (server *MainServer) Start(Db *sql.DB) {
	err := db.DatabaseInit(Db)
	if err != nil {
		log.Fatal("Can't init database err = ", err)
	}
	server.InitRoutes()
}
func (server *MainServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	server.router.ServeHTTP(writer, request)
}
