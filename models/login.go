package models

type Login struct {
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}
