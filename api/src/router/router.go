package router

import (
	routes "api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate creates a new router
func Generate() *mux.Router {
	r := mux.NewRouter()

	return routes.Config(r)
}
