package main

import (
	"fmt"
	"sync"
)

func count1(number int, how_much int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < how_much; i++ {
		number = number + number
		fmt.Println(number)
	}
}

func count2(number int, how_much int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < how_much; i++ {
		number += number
		fmt.Println(number)
	}
}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	number1 := 3
	number2 := 5

	how_much := 10

	go count1(number1, how_much, &wg)
	go count2(number2, how_much, &wg)

	wg.Wait()
}
