package main

import (
	"fmt"
)

func reverseArray(arr [8]string) {

	length := len(arr)

	for i := 0; i < length/2; i++ {
		arr[i], arr[length-1-i] = arr[length-1-i], arr[i]
	}

	fmt.Println(arr)
}

func main() {
	arr := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}

	reverseArray(arr)
}
