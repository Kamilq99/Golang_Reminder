package main

import (
	"fmt"
)

func main() {

	m := 4
	n := 6

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Print("*")
		fmt.Println(" ")
	}
}
