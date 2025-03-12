package main

import "fmt"

type Human struct {
	name string
	age  int
}

func updateData(h1 *Human) {
	h1.name = "Kamil"
	h1.age = 25
}

func main() {
	h := Human{name: "John", age: 30}

	fmt.Println(h)

	updateData(&h)

	fmt.Println(h)
}
