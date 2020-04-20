package rourter

import (
	hconfig "CrytoServer/client"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Route() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/symbol/{symbol}", hconfig.GetSymbol).Methods("GET")
	router.HandleFunc("/currency/all", hconfig.GetAll).Methods("GET")

	log.Fatal(http.ListenAndServe(":8090", router))

}
