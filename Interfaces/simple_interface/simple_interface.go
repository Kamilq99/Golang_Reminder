package main

import (
	"fmt"
)

type Cat struct {
	Name string
}

type Speaker interface {
	Speak() string
}

func (c *Cat) Speak() string {
	fmt.Printf("%s says ", c.Name)
	return "meow"
}

type Dog struct {
	Name string
}

func (d *Dog) Speak() string {
	fmt.Printf("%s says ", d.Name)
	return "woof"
}

func main() {

	c := Cat{Name: "Whiskers"}
	s := Speaker(&c)
	fmt.Println(s.Speak())

	d := Dog{Name: "Fido"}
	s = Speaker(&d)
	fmt.Println(s.Speak())

}
