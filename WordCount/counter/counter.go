package counter

import "strings"

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)

	count := make(map[string]int)

	for _, elem := range fields {
		count[elem] += 1
	}
	return count
}
