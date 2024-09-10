package service

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	var router *mux.Router = mux.NewRouter()

	for _, rou := range routes {
		router.Methods(rou.Method).
			Name(rou.Name).
			Path(rou.Pattern).
			Handler(rou.HandlerFunc)
	}
	return router
}
