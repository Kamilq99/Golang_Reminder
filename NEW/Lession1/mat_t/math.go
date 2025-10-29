package main

import (
	"fmt"
)

func main() {

	var a int
	var b int

	var subtract int
	var divide int

	fmt.Print("Enter number a = ")
	fmt.Scan(&a)
	fmt.Print("Enter number b = ")
	fmt.Scan(&b)

	add := a + b

	if a < b {
		fmt.Println("a is smaller than b")
	} else {
		subtract = a - b
	}

	if b == 0 {
		fmt.Println("b is zero, you cannot divide by zero")
	} else {
		divide = a / b
	}

	multiply := a * b

	fmt.Println("a + b = ", add)
	fmt.Println("a - b = ", subtract)
	fmt.Println("a / b = ", divide)
	fmt.Println("a * b = ", multiply)
}
