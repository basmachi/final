package app

import (
	"final/models"
	"fmt"
	"log"
	"net/http"
)

func (server *MainServer) InitRoutes(){
	fmt.Println("Server is listening in localhost:8888")
	//TODO ROUTES
	test(server)
//	server.router.POST("/api/login", server.)
	server.router.POST("/api/add-user", server.AddNewStudentHandler)
	err := http.ListenAndServe(":8888", server)
	if err != nil {
		log.Fatalf("err is %e", err)
	}
}

func test(server *MainServer)  {
	var a models.User
	a.Login = "surush"
	a.Password = "suurhs"
	a.Name = "surush"
	a.Surname = "surush"
	a.Age = "16"
	a.Phone = 48454
	a.Gender = "man"
	a.Role ="admin"
	a.Email = "aaa@mail.com"
	err := server.svc.AddNewUser(a)
	if err != nil {
		fmt.Println(err)
	}
}