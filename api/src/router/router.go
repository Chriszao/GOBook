package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Returns a router with configured routes
func Generate() *mux.Router {
	router := mux.NewRouter()

	return routes.Config(router)
}
