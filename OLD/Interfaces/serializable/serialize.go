package main

import "fmt"

type User struct {
	Name string
}

type Product struct {
	Product_name string
}

type Serializable interface {
	ToJSON() string
}

func (u User) ToJSON() string {
	return "JSON representation of User"
}

func (p Product) ToJSON() string {
	return "JSON representation of Product"
}

func main() {

	u := User{Name: "John"}
	p := Product{Product_name: "Laptop"}

	u.ToJSON()
	p.ToJSON()

	fmt.Printf("User: %s \n", u.ToJSON())
	fmt.Printf("Product: %s \n", p.ToJSON())
}
