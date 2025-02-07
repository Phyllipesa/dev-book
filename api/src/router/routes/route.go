package routes

import (
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

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
