package service

import (
	"encoding/json"
	"goblog/accountservice/dbclient"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var DbClient dbclient.IBoltClient

func Hello(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("hello"))
}

// 获取账号
func GetAccount(w http.ResponseWriter, req *http.Request) {
	//id := req.URL.Query().Get("accountId")
	id := mux.Vars(req)["accountId"]
	log.Println("id =", id)
	account, err := DbClient.QueryAccount(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(err.Error()))
		return
	}

	data, err := json.Marshal(account)
	if err != nil {
		log.Fatalf("parse account fail %s.", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func Test(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL.Path)

	w.Header().Set("Content-Type", "taxt/plain")
	w.Write([]byte("test."))

}
