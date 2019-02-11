package router

import (
	"github.com/anthonydenecheau/gopocservice/person"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreateRoutesGeneric(router *mux.Router) *mux.Router {

	// endpoints attachés au routeur
	for _, route := range routes {
		log.Println("Add route {}", route.Name)
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
func CreateRoutesApi(router *mux.Router) *mux.Router {

	// endpoints attachés aux API
	for _, route := range person.NewRoutes() {
		log.Println("Add route {}", route.Name)
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
func NewRouter() *mux.Router {

	log.Println("Create router ... ")

	router := mux.NewRouter().
		StrictSlash(true)

	// endpoints attachés au routeur
	CreateRoutesGeneric(router)

	// endpoints attachés aux API
	r := router.
		StrictSlash(true).
		PathPrefix("/v1/people").Subrouter()
	CreateRoutesApi(r)

	// endpoints api secure
	amw := authenticationMiddleware{make(map[string]string)}
	amw.Populate()
	r.Use(amw.Middleware)

	return router

}
