package main

import "fmt"

func main() {

	var fl float32
	fmt.Println("Enter a float number")

	num, err := fmt.Scan(&fl)

	fmt.Println(int(fl))

	fmt.Println(num)
	fmt.Println(err)

}
