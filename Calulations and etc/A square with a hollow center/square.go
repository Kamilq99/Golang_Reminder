package main

import (
	"fmt"
)

/*
func main() {
	n := 4

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 || i == n {
				fmt.Print("*")
			} else {
				if j == 1 || j == n {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println(" ")
	}
}
*/

func main() {
	n := 4

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 || i == n || j == 1 || j == n {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
	}
}
