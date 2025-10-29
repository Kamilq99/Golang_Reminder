package main

import "fmt"

type Student struct {
	name    string
	surname string
	grades  []float64
}

func getNewGrade(student *Student, newGrade float64) {

	student.grades = append(student.grades, newGrade)
}

func getAverage(student *Student) float64 {

	var total float64

	for _, grade := range student.grades {
		total += grade
	}

	return total / float64(len(student.grades))
}

func main() {

	student := Student{name: "Jan", surname: "Kowalski", grades: []float64{4.5, 4.5, 5.0, 5.0, 4.0}}

	fmt.Println(student)

	getNewGrade(&student, 5.0)

	fmt.Println(student)

	fmt.Println(getAverage(&student))
}
