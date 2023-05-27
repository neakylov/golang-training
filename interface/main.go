package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Square struct {
	sideLength float64
}

func (s Square) Area() float64 {
	return s.sideLength * s.sideLength
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	square := Square{10}
	printShapeArea(square)
	circle := Circle{5}
	printShapeArea(circle)
}

func printShapeArea(s Shape) {
	fmt.Println(s.Area())
}
