package main

import (
	"testing"
)

type isLetterTest struct {
	given    string
	expected bool
}

type isNumberTest struct {
	given    string
	expected bool
}

func TestIsLetter(t *testing.T) {

	var isLetterTests = []isLetterTest{
		{"a", true},
		{"b", true},
		{"c", true},
		{"5", false},
		{"!", false},
	}

	for _, test := range isLetterTests {
		if got := isLetter(test.given); got != test.expected {
			t.Errorf("Incorrect value returned on input: %s  | got: %t, want: %t", test.given, got, test.expected)
		}
	}

}

func TestIsNumber(t *testing.T) {

	var isNumberTests = []isNumberTest{
		{"1", true},
		{"2", true},
		{"3", true},
		{"f", false},
		{"!", false},
	}

	for _, test := range isNumberTests {
		if got := isNumber(test.given); got != test.expected {
			t.Errorf("Incorrect value returned on input: %s  | got: %t, want: %t", test.given, got, test.expected)
		}
	}
}
