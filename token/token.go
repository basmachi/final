package token

import (
	"final/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type MyCustomClaims struct {
	ID    int64
	Role  string
	Login string
	jwt.StandardClaims
}

func CreateToken(user models.User) string {
	mySigningKey := []byte("AllYourBase")
	claims := MyCustomClaims{
		ID:    user.ID,
		Role:  user.Role,
		Login: user.Login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 3000000,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	totalToken, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Println(err)

	}
	fmt.Printf("I am a token = %v\n", totalToken)
	return totalToken
}

func ParseToken(tokenString string) *MyCustomClaims {
	token, _ := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})
	claims := token.Claims.(*MyCustomClaims)
	return claims
}
