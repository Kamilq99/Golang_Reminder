package main

import "fmt"

type Person struct {
	name    string
	age     int
	address Address
}

type Address struct {
	city   string
	street string
}

func ChangeAddress(p *Person) {
	p.address.city = "Cape Town"
	p.address.street = "Long Street"
}

func main() {

	p := Person{
		name: "John",
		age:  30,
		address: Address{
			city:   "New York",
			street: "Wall Street",
		},
	}

	fmt.Println(p)

	ChangeAddress(&p)

	fmt.Println(p)

}
