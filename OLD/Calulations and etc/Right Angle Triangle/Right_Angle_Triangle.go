package main

import (
	"fmt"
)

func main() {

	n := 10

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j <= i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
