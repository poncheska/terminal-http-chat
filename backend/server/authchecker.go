package server

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"terminal-http-chat/server/database"
	"terminal-http-chat/server/utils"
)


type AuthData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func AuthChecker(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := jwt.Parse(utils.AuthKey, func(token *jwt.Token) (interface{}, error) {
			return utils.AuthKey, nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var newReq *http.Request
		if token.Valid {
			b := database.CheckUserData(token.Claims["login"], token.Claims)
			if !b {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			newReq = r.WithContext(context.WithValue(r.Context(), "user", map[string]string{
				"login":    claims.Login,
				"password": claims.Password,
			}))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		handlerFunc(w, newReq)
	}
}
