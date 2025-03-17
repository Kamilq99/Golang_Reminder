package main

type Car struct {
	Brand string
	Model string
	Year  int
}

type Bike struct {
	Brand string
	Model string
	Year  int
}

type Vehicle interface {
	Drive()
	Stop()
}

func (c Car) Drive() {
	println("Drive Car")
}
func (c Car) Stop() {
	println("Stop Car")
}
func (b Bike) Drive() {
	println("Drive Bike")
}

func (b Bike) Stop() {
	println("Stop Bike")
}

func main() {

	c := Car{Brand: "Ford", Model: "Mustang", Year: 2020}
	b := Bike{Brand: "Honda", Model: "CBR", Year: 2021}

	var v Vehicle

	v = c
	v.Drive()
	v.Stop()

	v = b
	v.Drive()
	v.Stop()
}
