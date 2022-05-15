package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//  All API routes must follow this format
type Route struct {
	URI       string
	Method    string
	Function  func(http.ResponseWriter, *http.Request)
	IsPrivate bool
}

// Put all routes on the router
func Config(router *mux.Router) *mux.Router {
	routes := userRoutes

	routes = append(routes, loginRoute)

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return router
}
