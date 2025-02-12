package main

import (
	"fmt"
)

func reverseString(text string) string {

	runes := []rune(text)
	length := len(runes)

	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}
	return string(runes)
}

func main() {

	text := "Hello World!"

	newText := reverseString(text)

	fmt.Println(newText)
}
