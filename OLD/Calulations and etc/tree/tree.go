package main

import (
	"fmt"
)

func main() {
	height := 3 // Możesz zmienić wysokość choinki

	for i := 0; i < height; i++ {
		// Drukowanie spacji
		for j := 0; j < height-i-1; j++ {
			fmt.Print(" ")
		}
		// Drukowanie gwiazdek
		for k := 0; k < 2*i+1; k++ {
			fmt.Print("*")
		}
		fmt.Println() // Przejście do nowej linii
	}
}
