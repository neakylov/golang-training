package NotVerySecure

import "testing"

func TestAlphanumeric(t *testing.T) {
	testTable := []struct {
		str      string
		expected bool
	}{
		{
			str:      ".*?",
			expected: false,
		},
		{
			str:      "a",
			expected: true,
		},
		{
			str:      "Mazinkaiser",
			expected: true,
		},
		{
			str:      "hello world_",
			expected: false,
		},
		{
			str:      "PassW0rd",
			expected: true,
		},
		{
			str:      "     ",
			expected: false,
		},
		{
			str:      "",
			expected: false,
		},
		{
			str:      "\n\t\n",
			expected: false,
		},
		{
			str:      "ciao\n$$_",
			expected: false,
		},
		{
			str:      "__ * __",
			expected: false,
		},
		{
			str:      "&)))(((",
			expected: false,
		},
		{
			str:      "43534h56jmTHHF3k",
			expected: true,
		},
	}

	for _, testCase := range testTable {
		result := alphanumeric(testCase.str)

		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %v, got %v", testCase.expected, result)
		}
	}
}
