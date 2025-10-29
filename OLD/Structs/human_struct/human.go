package main

import (
	"fmt"
)

type human struct {
	name string
	age  int
}

func sort(arr []human) {

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i].age > arr[j].age {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func main() {
	h := human{name: "John", age: 30}

	fmt.Println(h)

	arr := [5]human{{name: "John", age: 30}, {name: "Jane", age: 25}, {name: "Kamil", age: 26}, {name: "Robert", age: 27}, {name: "Bob", age: 19}}

	fmt.Println(arr)

	sort(arr[:])

	fmt.Println(arr)
}
