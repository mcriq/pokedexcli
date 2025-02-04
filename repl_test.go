package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "learning go is awesome",
			expected: []string{"learning", "go", "is", "awesome"},
		},
		{
			input: "hello",
			expected: []string{"hello"},
		},
		{
			input: "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		if len(actual) != len(c.expected) {
			t.Errorf("slices do not match. Expected: %v\n Actual: %v", c.expected, actual)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("actual word: %s does not match expected word: %s", word, expectedWord)
			}
		}
	}
}