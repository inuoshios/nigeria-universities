package router

import (
	"github.com/gorilla/mux"
)

// NEW creates a new router.
func NEW() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return Route(r)
}
