package router

import (
	conf "GoWeb1/config"
	hconfig "GoWeb1/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", hconfig.HomeLink)
	router.HandleFunc("/user", hconfig.CreateEvent).Methods("POST")
	router.HandleFunc("/doc/{id}/{name}", hconfig.GetEvent).Methods("GET")
	router.HandleFunc("/docs", hconfig.GetEvents).Methods("GET")
	router.HandleFunc("/update", hconfig.UpdateEvent).Methods("PUT")

	log.Fatal(http.ListenAndServe(conf.Config.PortNo, router))
}
