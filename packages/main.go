package main

import (
	"fmt"

	"packages/greetings"
)

func main() {
	message := greetings.SayHello("from GO")
	fmt.Println(message)
}
