package main

import (
	"fmt"
)

func countChar(text string, letter string) int {

	runes := []rune(text)
	length := len(runes)
	count := 0

	for i := 0; i < length; i++ {
		if string(runes[i]) == letter {
			count++
		}
	}
	return count
}

func main() {

	var letter string
	fmt.Print("Get a letter: ")
	fmt.Scan(&letter)

	text := "Go is a simple and powerful programming language."

	result := countChar(text, letter)

	fmt.Printf("result = %d", result)
}
