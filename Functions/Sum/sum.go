package main

import (
	"fmt"
)

func sum(a int, b int) int {

	result := a + b

	return result
}

func main() {
	var a int
	var b int

	fmt.Print("Enter number a = ")
	fmt.Scan(&a)
	fmt.Print("Enter number b = ")
	fmt.Scan(&b)

	result := sum(a, b)

	fmt.Printf("Result = %d", result)
}
