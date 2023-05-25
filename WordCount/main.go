package main

import (
	"fmt"
	"wordcount/counter"
)

func main() {
	s := "I love this language!"
	fmt.Println(counter.WordCount(s))
}
