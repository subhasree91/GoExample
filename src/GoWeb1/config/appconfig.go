package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var file = flag.String("f", "defaultconfig.json", "filepath")
var Config Appconfig

func ReadConfig() {
	fmt.Println("filepath:", *file)
	flag.Parse()
	// fileout, err := os.Open(*file)
	// if err != nil {
	// 	fmt.Println("file error", err)
	// 	os.Exit(1)
	// }

	// defer fileout.Close()

	arr, err := ioutil.ReadFile(*file)

	if err != nil {
		fmt.Println("Error in reading the file:", err)
		os.Exit(1)
	}
	//arr, err := json.Marshal(fileout)
	// fmt.Println("arr:", string(arr))
	// if err != nil {
	// 	fmt.Println("Config error")
	// 	os.Exit(1)
	// }

	err = json.Unmarshal(arr, &Config)

	// decoder := json.NewDecoder(fileout)

	// err = decoder.Decode(&Config)

	if err != nil {
		fmt.Println("Error in unmarshalling:", err)
		os.Exit(1)
	}
	fmt.Printf("configin app %v\n", Config)

}
