package arrays

import "testing"

func TestNewArray(t *testing.T) {
	testTable := []struct {
		cap       int
		maxNumber int
	}{
		{
			cap:       10,
			maxNumber: 10,
		},
		{
			cap:       15,
			maxNumber: 32,
		},
		{
			cap:       3,
			maxNumber: 100,
		},
		{
			cap:       20,
			maxNumber: 1000,
		},
		{
			cap:       10,
			maxNumber: 11,
		},
	}

	for _, testCase := range testTable {
		array := NewArray(testCase.cap, testCase.maxNumber)
		if len(array) != testCase.cap || cap(array) != testCase.cap {
			t.Errorf("Incorrect result. Expect len: %d, cap: %d. Got len: %d, cap: %d", testCase.cap, testCase.cap, len(array), cap(array))
		}

		for _, elem := range array {
			if elem > testCase.maxNumber {
				t.Errorf("Incorrect result. %d > %d", elem, testCase.maxNumber)
			}
		}
	}
}

func TestMaxValue(t *testing.T) {
	testTable := []struct {
		array    []int
		expected int
	}{
		{
			array: []int{
				32, 56, 23, 65, 67,
			},
			expected: 67,
		},
		{
			array: []int{
				76, 21, 65, 33, 44, 87, 99, 777, 1, 11,
			},
			expected: 777,
		},
		{
			array: []int{
				1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
			},
			expected: 10,
		},
		{
			array: []int{
				14, 46, 49, 97, 90, 28, 64, 26, 27, 71,
			},
			expected: 97,
		},
		{
			array: []int{
				18, 35, 75, 9, 10, 83, 91, 38, 99, 100,
			},
			expected: 100,
		},
	}

	for _, testCase := range testTable {
		value := MaxValue(testCase.array)

		if value != testCase.expected {
			t.Errorf("Incorrect result. Expect %d, got %d", testCase.expected, value)
		}
	}
}
