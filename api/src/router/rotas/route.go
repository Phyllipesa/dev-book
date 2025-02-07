package rotas

import "net/http"

// Route represents a route in the application
type Route struct {
	URI                   string
	Method                string
	Function              func(w http.ResponseWriter, r *http.Request)
	RequestAuthentication bool
}
