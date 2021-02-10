package utils

import "github.com/dgrijalva/jwt-go"

var (
	AuthKey = "VeryStronkkkKey"
)

func CreateToken(login, password string) string{
	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims["login"] = login
	token.Claims["password"] = password
}