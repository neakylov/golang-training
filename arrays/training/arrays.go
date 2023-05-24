package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var a [10]int

	for index := range a {
		a[index] = rand.Intn(1000)
	}

	var max int

	for _, elem := range a {
		if elem > max {
			max = elem
		}
	}
	fmt.Println(max)
}
