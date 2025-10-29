package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Generator interface {
	Generate() any
}

type RandomInt struct{}

func (r RandomInt) Generate() any {
	return rand.Intn(100)
}

type RandomString struct{}

func (r RandomString) Generate() any {
	letters := "abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, 5)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var generators = []Generator{
		RandomInt{},
		RandomString{},
	}

	for _, gen := range generators {
		fmt.Println(gen.Generate())
	}
}
