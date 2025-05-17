package main

import (
	"testing"
)

func TestGetBFCode(t *testing.T) {
	// Test file io
	var expected string = "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."
	var actual string = getBFCode("../Examples/helloWorld.bf")

	if expected != actual {
		t.Errorf("Expected: %s\n Actual: %s", expected, actual)
	}
}

func TestTokenize(t *testing.T) {
	codeToTokenize := "><+-.,[]"
	expected := []Token{{">", 1, 1}, {"<", 1, 2}, {"+", 1, 3}, {"-", 1, 4},
		{".", 1, 5}, {",", 1, 6}, {"[", 1, 7}, {"]", 1, 8}}
	actual := tokenize(codeToTokenize)

	// Check that the lengths of the slices are equal
	if len(expected) != len(actual) {
		t.Errorf("Length of resulting slice does not match length of the expected slice")
		return
	}

	// Loop through each index and check that the
	for i, _ := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Token at index %d does not match. Expected {%s, %d, %d}. Actual {%s, %d, %d}",
				i, expected[i].token, expected[i].row, expected[i].col, actual[i].token, actual[i].row, actual[i].col)
		}
	}
}

func TestTokenizeWithCharacters(t *testing.T) {
	codeToTokenize := "><+-abcd.,[]This code should not be tokenized because it is a comment"
	expected := []Token{{">", 1, 1}, {"<", 1, 2}, {"+", 1, 3}, {"-", 1, 4},
		{".", 1, 9}, {",", 1, 10}, {"[", 1, 11}, {"]", 1, 12}}
	actual := tokenize(codeToTokenize)

	// Check that the lengths of the slices are equal
	if len(expected) != len(actual) {
		t.Errorf("Length of resulting slice does not match length of the expected slice")
		return
	}

	// Loop through each index and check that the
	for i, _ := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Token at index %d does not match. Expected {%s, %d, %d}. Actual {%s, %d, %d}",
				i, expected[i].token, expected[i].row, expected[i].col, actual[i].token, actual[i].row, actual[i].col)
		}
	}
}

func TestTokenizeWithNewLines(t *testing.T) {
	codeToTokenize := "><+-abcd.,[]\nThis code should not be tokenized because it is a comment\n><.,"
	expected := []Token{{">", 1, 1}, {"<", 1, 2}, {"+", 1, 3}, {"-", 1, 4},
		{".", 1, 9}, {",", 1, 10}, {"[", 1, 11}, {"]", 1, 12},
		{">", 3, 1}, {"<", 3, 2}, {".", 3, 3}, {",", 3, 4}}
	actual := tokenize(codeToTokenize)

	// Check that the lengths of the slices are equal
	if len(expected) != len(actual) {
		t.Errorf("Length of resulting slice does not match length of the expected slice")
		return
	}

	// Loop through each index and check that the
	for i, _ := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Token at index %d does not match. Expected {%s, %d, %d}. Actual {%s, %d, %d}",
				i, expected[i].token, expected[i].row, expected[i].col, actual[i].token, actual[i].row, actual[i].col)
		}
	}
}
