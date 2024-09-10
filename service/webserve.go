package service

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartWebService(port string, appName string) {

	var mux *mux.Router = NewRouter()
	log.Println("Started HTTP service " + appName + " on: " + port)
	log.Fatalln(http.ListenAndServe(":"+port, mux))

}
