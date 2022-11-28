package router

import (
	"github.com/gorilla/mux"
)

// NEW creates a new router.
func NEW() *mux.Router {
	r := mux.NewRouter()
	return Route(r)
}
