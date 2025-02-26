package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents a route in the application
type Route struct {
	URI                   string
	Method                string
	Function              func(w http.ResponseWriter, r *http.Request)
	RequestAuthentication bool
}

// Config receives all routes and returns a router with all routes configured
func Config(r *mux.Router) *mux.Router {
	routes := routeUsers
	routes = append(routes, routeLogin)
	routes = append(routes, routePublications...)

	for _, route := range routes {
		if route.RequestAuthentication {
			r.HandleFunc(
				route.URI,
				middlewares.Logger(middlewares.Authentication(route.Function)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(
				route.URI,
				middlewares.Logger(route.Function),
			).Methods(route.Method)
		}
	}

	return r
}
