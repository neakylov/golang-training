package arrays

import (
	"math/rand"
)

func NewArray(cap int, maxNumber int) []int {
	array := make([]int, cap)

	for i := 0; i < cap; i++ {
		array[i] = rand.Intn(maxNumber)
	}
	return array
}

func MaxValue(array []int) int {
	var max int

	for _, elem := range array {
		if elem > max {
			max = elem
		}
	}
	return max
}
