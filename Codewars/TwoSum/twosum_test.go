package TwoSum

import "testing"

func TestTwoSum(t *testing.T) {
	testTable := []struct {
		numbers  []int
		target   int
		expected [2]int
	}{
		{
			numbers:  []int{1, 2, 3},
			target:   4,
			expected: [2]int{0, 2},
		},
		{
			numbers:  []int{1234, 5678, 9012},
			target:   14690,
			expected: [2]int{1, 2},
		},
		{
			numbers:  []int{2, 2, 3},
			target:   4,
			expected: [2]int{0, 1},
		},
	}
	for _, testCase := range testTable {
		result := TwoSum(testCase.numbers, testCase.target)

		for index := range result {
			if result[index] != testCase.expected[index] {
				t.Errorf("Incorrect result. Expect %v, got %v", testCase.expected, result)
			}
		}
	}
}
