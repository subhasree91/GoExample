package main

import "fmt"

func missingNo() int {

	arr := [6]int{1, 2, 3, 5, 6, 7}

	var n int = len(arr)
	lastNo := arr[n-1]

	var exSum int = (lastNo * (lastNo + 1)) / 2
	var accSum int = 0

	var missNo int = 0

	for i := 0; i <= len(arr)-1; i++ {

		accSum = accSum + arr[i]
	}

	if (exSum - accSum) != 0 {
		missNo = (exSum - accSum)
	}

	return missNo

}

func main() {

	sol := missingNo()
	fmt.Println(sol)

}
