package models

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Age string `json:"age"`
	Phone    int64  `json:"phone"`
	Email    string `json:"email"`
	Gender	 string `json:"gender"`
	Role     string `json:"role"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Remove   bool   `json:"remove"`
}