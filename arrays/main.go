package main

import (
	"arrays/arrays"
	"fmt"
)

func main() {
	array := arrays.NewArray(15, 100)
	fmt.Println(arrays.MaxValue(array))
}
