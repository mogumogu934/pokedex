package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello world   ",
			expected: []string{"hello", "world"},
		},

		{
			input:    "  PikaChU is  cute! ",
			expected: []string{"pikachu", "is", "cute!"},
		},

		{
			input:    "     ",
			expected: []string{},
		},

		{
			input:    "mudkip",
			expected: []string{"mudkip"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("got %d words, want %d words", len(actual), len(c.expected))
		}

		for i := range actual {
			actualWord := actual[i]
			expectedWord := c.expected[i]
			if actualWord != expectedWord {
				t.Errorf("got %s, expected %s", actualWord, expectedWord)
			}
		}
	}
}
