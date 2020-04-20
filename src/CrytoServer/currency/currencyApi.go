package currency

import (
	"CrytoServer/utility"
	"encoding/json"
	"fmt"
)

func GetAll() ([]Currency, error) {

	var currencyArr []Currency

	tempArr, err := utility.HttpReq("GET", CurrencyAllUrl, nil)

	if err != nil {
		fmt.Println("Error in GetAll", err)
		return currencyArr, err
	}
	err = json.Unmarshal(tempArr, &currencyArr)
	if err != nil {
		fmt.Println("Error in Unmarshall GETAll", err)
	}
	return currencyArr, err

}

func GetSymbol(symbol string) (Symbol, error) {

	var symbolArr Symbol

	tempArr, err := utility.HttpReq("GET", CurencySymbolUrl+"/"+symbol, nil)

	if err != nil {
		fmt.Println("Error in GetSymbol", err)
		return symbolArr, err
	}
	err = json.Unmarshal(tempArr, &symbolArr)
	if err != nil {
		fmt.Println("Error in Unmarshall GETSymbol", err)
	}
	return symbolArr, err
}
