package main

import (
	"fmt"
	"strings"
)

func main() {

	var St string
	fmt.Println("Enter a String")

	num, err := fmt.Scan(&St)

	StringLow := strings.ToLower(St)

	if strings.IndexByte(St, 'i') == 0 {
		if (strings.IndexByte(St, 'n')) == (len(StringLow) - 1) {
			if (strings.IndexAny(St, "a")) != 0 {
				fmt.Println("Found")
			}

		}
	} else {
		fmt.Println("NotFound")
	}

	fmt.Println(num)
	fmt.Println(err)

}
