package middlewares

import (
	"api/src/config"
	"api/src/responses"
	"log"
	"net/http"
)

// Prints request into console
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("\n %s %s %s", request.Method, request.RequestURI, request.Host)

		next(writer, request)
	}
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if err := config.ValidateToken(request); err != nil {
			responses.Error(writer, http.StatusUnauthorized, err)
			return
		}
		next(writer, request)
	}
}
