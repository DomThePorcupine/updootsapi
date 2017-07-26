package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

/*
APIRouter creates and then exports a router
to serve the updoots API woo
*/
func APIRouter() *mux.Router {
	router := mux.NewRouter()

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		router.Methods(route.Method).Path("/api/v1" + route.Pattern).Name(route.Name).Handler(handler)
	}

	for _, route := range protectedroutes {
		var handler http.Handler

		handler = route.HandlerFunc
		router.Methods(route.Method).Path("/api/v1" + route.Pattern).Name(route.Name).HandlerFunc(validate(handler))
	}

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./pub/dist")))
	router.PathPrefix("/releases").Handler(http.FileServer(http.Dir("./releases")))
	http.Handle("/", router)

	return router
}
