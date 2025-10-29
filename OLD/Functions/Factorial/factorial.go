package main

import (
	"fmt"
)

func factorial(number int) int {

	result := 1

	for i := 1; i <= number; i++ {
		result *= i
	}

	return result
}

func main() {

	var number int

	fmt.Print("Get a number")
	fmt.Scan(&number)

	result := factorial(number)

	fmt.Printf("Result = %d", result)
}
