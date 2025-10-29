package main

import (
	"fmt"
)

func change_value(p *float64) {

	*p = 20
}

func main() {

	var s float64 = 10

	fmt.Println("Before Change", s)

	change_value(&s)

	fmt.Println("After Change: ", s)
}
