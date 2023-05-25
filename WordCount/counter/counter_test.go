package counter

import (
	"strings"
	"testing"
)

func TestWordCount(t *testing.T) {
	testTable := []struct {
		str      string
		expected map[string]int
	}{
		{
			str: "I am learning Go!",
			expected: map[string]int{
				"I": 1, "am": 1, "learning": 1, "Go!": 1,
			},
		},
		{
			str: "The quick brown fox jumped over the lazy dog.",
			expected: map[string]int{
				"The": 1, "quick": 1, "brown": 1, "fox": 1, "jumped": 1,
				"over": 1, "the": 1, "lazy": 1, "dog.": 1,
			},
		},
		{
			str: "I ate a donut. Then I ate another donut.",
			expected: map[string]int{
				"I": 2, "ate": 2, "a": 1, "donut.": 2, "Then": 1, "another": 1,
			},
		},
		{
			str: "A man a plan a canal panama.",
			expected: map[string]int{
				"A": 1, "man": 1, "a": 2, "plan": 1, "canal": 1, "panama.": 1,
			},
		},
	}

	for _, testCase := range testTable {
		result := WordCount(testCase.str)
		fields := strings.Fields(testCase.str)

		for _, elem := range fields {
			if testCase.expected[elem] != result[elem] {
				t.Errorf("Incorrect result. Expect %v, got %v", testCase.expected, result)
				break
			}
		}
	}
}
