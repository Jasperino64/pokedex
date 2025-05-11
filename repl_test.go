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
			input: "   hello   world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Go   is   awesome  ",
			expected: []string{"go", "is", "awesome"},
		},
		{
			input:    "  Gopher   ",
			expected: []string{"gopher"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected %d words, got %d", len(c.expected), len(actual))
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("Expected %s, got %s", c.expected[i], actual[i])
			}
		}
	}
}