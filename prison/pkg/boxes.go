package pkg

import (
	"math/rand"
)

type Box struct {
	number int
}

func GetBoxes() [100]Box {
	var boxes [100]Box

	for index := range boxes {
		boxes[index].number = index
	}

	rand.Shuffle(100, func(i, j int) {
		boxes[i].number, boxes[j].number = boxes[j].number, boxes[i].number
	})

	return boxes
}

func CheckBox(person Person, boxes [100]Box) bool {
	const attempts = 50
	boxNumber := boxes[person.number].number

	for attempt := 0; attempt < attempts; attempt++ {
		if person.number == boxNumber {
			return true
		}
		boxNumber = boxes[boxNumber].number

	}
	return false
}
