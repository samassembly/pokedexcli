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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input: "CHANSEY WhOOper ditto",
			expected: []string{"chansey", "whooper", "ditto"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		if len(actual) != len(c.expected) {
			// if they don't match, use t.Errorf to print an error message
			t.Errorf("Failure: Unexpected result length")
			// and fail the test
			t.Fail()
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			if word != expectedWord {
			// if they don't match, use t.Errorf to print an error message
			t.Errorf("Failure: Generated output does not match expected output")
			// and fail the test
			t.Fail()
			}
		}
	}
}