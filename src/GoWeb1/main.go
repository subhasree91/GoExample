package main

import (
	conf "GoWeb1/config"
	route "GoWeb1/router"
	"fmt"
)

// func homeLink(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome home!")
// }

func main() {

	conf.ReadConfig()
	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/", homeLink)
	// log.Fatal(http.ListenAndServe(conf.Config.PortNo, router))

	fmt.Printf("config %v\n", conf.Config)

	route.Router()

}
