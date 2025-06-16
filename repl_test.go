package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " Hello World   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Bla nah   ",
			expected: []string{"bla", "nah"},
		},
		{
			input:    "   ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Failed expected length of array check")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Failed expected word check. Expected %s, got %s.", expectedWord, word)
			}
		}
	}
}
