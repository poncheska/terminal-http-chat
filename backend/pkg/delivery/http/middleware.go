package http

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	authHeader   = "Authorization"
	userIdHeader = "UserId"
)

func (h Handler) AuthChecker(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authHeader)

		if header == "" {
			WriteUnauthorizedResponse(w, "auth header is empty")
			return
		}

		headerParts := strings.Split(header, " ")

		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			WriteUnauthorizedResponse(w, "invalid auth header")
			return
		}

		if headerParts[1] == "" {
			WriteUnauthorizedResponse(w, "token is empty")
			return
		}

		id, err := h.tokenService.ParseToken(headerParts[1])
		if err != nil {
			WriteUnauthorizedResponse(w, err.Error())
			return
		}

		log.Printf("user: %v authorized", id)

		r.Header.Set(userIdHeader, strconv.FormatInt(id, 10))
		handlerFunc(w, r)
	}
}
