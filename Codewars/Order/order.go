package Order

import (
	"strings"
	"unicode"
)

func Order(sentence string) string {
	if sentence == "" {
		return sentence
	}
	fields := strings.Fields(sentence)
	result := make([]string, len(fields))
	nums := map[string]int{
		"1": 0, "2": 1, "3": 2, "4": 3, "5": 4, "6": 5,
		"7": 6, "8": 7, "9": 8,
	}
	for _, elem := range fields {
		num := findNumber(elem)
		result[nums[num]] = elem
	}

	return strings.Join(result, " ")

}

func findNumber(s string) string {
	for _, elem := range s {
		if unicode.IsDigit(elem) {
			return string(elem)
		}
	}

	return ""
}
