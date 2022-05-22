package routes

import (
	"api/src/middlewares"
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

		if route.IsPrivate {
			router.HandleFunc(
				route.URI,
				middlewares.Logger(middlewares.Auth(route.Function)),
			).Methods(route.Method)
		} else {
			router.HandleFunc(
				route.URI,
				middlewares.Logger(route.Function),
			).Methods(route.Method)
		}
	}

	return router
}
