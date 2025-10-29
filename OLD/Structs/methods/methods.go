package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h Human) greet() {
	fmt.Println("Hello, my name is", h.name)
}

func (h Human) isAdult() bool {
	if h.age >= 18 {
		return true
	}
	return false
}

func main() {
	h := Human{name: "Kamil", age: 25}

	h.greet()

	fmt.Println("Is he an adult?", h.isAdult())
}
