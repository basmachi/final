package app

import (
	"encoding/json"
	"final/models"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

//Login
func (server *MainServer) LoginHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	//GET AND PARSE FROM FRONT
	var requestBody models.User
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	//CHECK ERROR
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode("json-invalid")
		if err != nil {
			log.Println("Can't find connection")
			return
		}
		return
	}

	err = server.svc.AddNewUser(requestBody)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(writer).Encode("err_server")
		if err != nil {
			log.Println("Can't find connection, err is ", err)
			return
		}
		return
	}
}

func (server *MainServer) AddNewStudentHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	//GET AND PARSE FROM FRONT
	var requestBody models.User
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	//CHECK ERROR
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode("json-invalid")
		if err != nil {
			log.Println("Can't find connection")
			return
		}
		return
	}

	err = server.svc.AddNewUser(requestBody)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(writer).Encode("err_server")
		if err != nil {
			log.Println("Can't find connection, err is ", err)
			return
		}
		return
	}
}
