package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ixxiv/nigeria-uni/internal/controllers"
	"github.com/ixxiv/nigeria-uni/internal/middlewares"
)

type Routes struct {
	URI     string
	Method  string
	Handler func(w http.ResponseWriter, r *http.Request)
}

var routes = []Routes{
	{
		URI:     "/",
		Method:  http.MethodGet,
		Handler: controllers.Welcome,
	},
	{
		URI:     "/v1",
		Method:  http.MethodGet,
		Handler: controllers.GetAllUniveristy,
	},
	{
		URI:     "/v1/{abbreviation}",
		Method:  http.MethodGet,
		Handler: controllers.GetSpecificUniversity,
	},
}

func Route(r *mux.Router) *mux.Router {
	for _, route := range routes {
		r.Use(middlewares.AddContentType)

		// adding PathPrefix
		routes := r.PathPrefix("/api").Subrouter()
		routes.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}

	return r
}
