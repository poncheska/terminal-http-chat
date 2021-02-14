package http

import (
	"net/http"
)

var (
	authHeader = "Authorization"
	userIdHeader = "UserId"
)


func AuthChecker(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authHeader)

		if header == ""{
			WriteErrorResponse(w, err)
			return
		}
		//token, err := jwt.Parse(utils.AuthKey, func(token *jwt.Token) (interface{}, error) {
		//	return utils.AuthKey, nil
		//})
		//if err != nil {
		//	w.WriteHeader(http.StatusUnauthorized)
		//	return
		//}
		//
		//var newReq *http.Request
		//if token.Valid {
		//	b := store.CheckUserData(token.Claims["login"], token.Claims)
		//	if !b {
		//		w.WriteHeader(http.StatusUnauthorized)
		//		return
		//	}
		//	newReq = r.WithContext(context.WithValue(r.Context(), "user", map[string]string{
		//		"login":    claims.Login,
		//		"password": claims.Password,
		//	}))
		//} else {
		//	w.WriteHeader(http.StatusUnauthorized)
		//	return
		//}
		//handlerFunc(w, newReq)
	}
}
