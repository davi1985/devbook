package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                   string
	Method                string
	Func                  func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}

func New() *mux.Router {
	r := mux.NewRouter()
	return configRoutes(r)
}

func configRoutes(r *mux.Router) *mux.Router {
	var routes []Route
	routes = append(routes, userRoutes...)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	return r
}
