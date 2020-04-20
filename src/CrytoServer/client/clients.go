package client

import (
	"CrytoServer/currency"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAll(w http.ResponseWriter, r *http.Request) {

	outputArr, err := currency.GetAll()

	if err != nil {
		fmt.Println("Error in client GetAll")

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode("Error in client GetAll")

		return
	}

	json.NewEncoder(w).Encode(outputArr)

}

func GetSymbol(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)

	outputArr, err := currency.GetSymbol(param["symbol"])

	if err != nil {
		fmt.Println("Error in client GetSymbol")

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode("Error in client GetSymbol")

		return
	}

	json.NewEncoder(w).Encode(outputArr)
}
