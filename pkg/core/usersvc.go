package core

import (
	"database/sql"
	"final/db"
	"final/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserSvc struct {
	db *sql.DB
}

func NewUserSvc(db *sql.DB) *UserSvc {
	if db == nil{
		log.Fatal(`Db is can't be nil`)
	}
	return &UserSvc{db: db}
}

func (receiver *UserSvc) AddNewUser(user models.User) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Can't do hash err is ", err)
		return err
	}
	user.Password = string(password)
	_, err = receiver.db.Exec(db.AddNewStudent, user.Name, user.Surname, user.Phone, user.Age, user.Email, user.Gender, user.Role, user.Login, user.Password)
	if err != nil {
		fmt.Println("Can't add new user")
		return err
	}
	return nil
}

//func (receiver *UserSvc) CheckLogin(user models.Login) error {
//	_, err := receiver.db.Exec("Select login")
//	if err != nil {
//		fmt.Println("Can't add new user")
//		return err
//	}
//	return nil
//}