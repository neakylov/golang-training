package Order

import "testing"

func TestOrder(t *testing.T) {
	testTable := []struct {
		sentence string
		expected string
	}{
		{
			sentence: "is2 Thi1s T4est 3a",
			expected: "Thi1s is2 3a T4est",
		},
		{
			sentence: "4of Fo1r pe6ople g3ood th5e the2",
			expected: "Fo1r the2 g3ood 4of th5e pe6ople",
		},
		{
			sentence: "",
			expected: "",
		},
	}
	for _, testCase := range testTable {
		result := Order(testCase.sentence)

		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %v, got %v", testCase.expected, result)
		}
	}
}

func TestFindNumber(t *testing.T) {
	testTable := []struct {
		word     string
		expected string
	}{
		{
			word:     "is2",
			expected: "2",
		},
		{
			word:     "Th1s",
			expected: "1",
		},
		{
			word:     "T4est",
			expected: "4",
		},
		{
			word:     "3a",
			expected: "3",
		},
	}

	for _, testCase := range testTable {
		result := findNumber(testCase.word)

		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %v, got %v", testCase.expected, result)
		}
	}
}
