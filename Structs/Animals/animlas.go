package main

import "fmt"

type Animal struct {
	Name    string
	Species string
	Sound   string
}

func (a Animal) Speak() {
	fmt.Printf("%s says %s \n", a.Name, a.Sound)
}

func AnimalsSpeak(animals []Animal) {
	for _, animal := range animals {
		fmt.Printf("%s says %s\n", animal.Name, animal.Sound)
	}
}

func main() {

	cat := Animal{Name: "Whiskers", Species: "Cat", Sound: "Meow"}

	cat.Speak()

	slice := []Animal{cat, {Name: "Fido", Species: "Dog", Sound: "Woof"}, {Name: "Rover", Species: "Dog", Sound: "Woof"}, {Name: "Mittens", Species: "Cat", Sound: "Meow"}}

	AnimalsSpeak(slice)
}
