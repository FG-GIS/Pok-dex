package main

import "testing"

func TestCleanInput(t *testing.T) {
	var failcount int = 0
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello  world     ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   jean  paul van  damme   ",
			expected: []string{"jean", "paul", "van", "damme"},
		},
		{
			input:    "   hablas  espanol?     ",
			expected: []string{"hablas", "espanol?"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}
	for i, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Test %v - Failed: wrong length on split", i)
			failcount++
		}
		for j := range actual {
			word := actual[j]
			testWord := c.expected[j]

			if word != testWord {
				t.Errorf("Test %v - Word %v wrong", i, j)
			}
		}
	}
}
