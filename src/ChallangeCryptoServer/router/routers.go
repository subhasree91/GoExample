package router

import (
	conf "ChallangeCryptoServer/config"
	hconfig "ChallangeCryptoServer/handler"
	"log"
	"net/http"
)

func router() {

	router.HandleFunc("/GET/currency/{symbol}", hconfig.GetCurrencySymbol).Methods("GET")
	router.HandleFunc("/GET/currency/all", hconfig.GetCurrenyAll).Methods("GET")

	log.Fatal(http.ListenAndServe(conf.Config.PortNo, router))

}
