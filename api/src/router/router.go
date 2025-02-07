package router

import "github.com/gorilla/mux"

// Generate creates a new router
func Generate() *mux.Router {
	return mux.NewRouter()
}
