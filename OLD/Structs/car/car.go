package main

import "fmt"

type Car struct {
	brand string
	model string
	year  int
}

func (c Car) Age() int {
	return 2025 - c.year
}

func main() {

	Mustang := Car{brand: "Ford", model: "Mustang", year: 2020}

	fmt.Println("Ford Mustang is", Mustang.Age(), "years old")

	Xsara := Car{brand: "Citroen", model: "Xsara", year: 2009}

	fmt.Println("Citroen Xsara is", Xsara.Age(), "years old")
}
