package main

import (
	"fmt"
	"prison/pkg"
)

func main() {
	prisoners := pkg.GetPrisoners()
	boxes := pkg.GetBoxes()

	for i, prisoner := range prisoners {
		result := pkg.CheckBox(prisoner, boxes)
		fmt.Printf("Prisoner number %v, result: %v \n", i, result)
	}
}
