package main

import (
	"fmt"
)

func isEven(number int) {

	if number%2 == 0 {
		fmt.Print("The Number is Even")
	} else {
		fmt.Print("The Number is Not Even")
	}
}

func main() {
	var number int

	fmt.Print("Get a number = ")
	fmt.Scan(&number)

	isEven(number)
}
