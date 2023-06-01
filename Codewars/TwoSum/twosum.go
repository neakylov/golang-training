package TwoSum

func TwoSum(numbers []int, target int) [2]int {
	var solution [2]int
	for currentIndex := 0; currentIndex < len(numbers); currentIndex++ {
		for index, elem := range numbers {
			if currentIndex == index {
				continue
			}
			if numbers[currentIndex]+elem == target {
				return [2]int{currentIndex, index}
			}
		}
	}
	return solution
}
