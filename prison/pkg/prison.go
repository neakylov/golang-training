package pkg

type Person struct {
	number int
}

func GetPrisoners() [100]Person {
	var prisoners [100]Person

	for index := range prisoners {
		prisoners[index].number = index
	}

	return prisoners
}
