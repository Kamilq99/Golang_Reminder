package main

import "fmt"

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Shape interface {
	Area() float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {

	c := Circle{Radius: 5}
	r := Rectangle{Width: 10, Height: 20}

	fmt.Println("Circle area:", c.Area())
	fmt.Println("Rectangle area:", r.Area())
}
