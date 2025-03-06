package main

import (
	"fmt"
)

func main() {

	a := 1

	b := &a

	fmt.Printf("Variable is equal to: %d \n", *b)
	fmt.Printf("Variable addres is equal to: %d", &a)
	fmt.Println("Pointer addres in b: ", b)
}
