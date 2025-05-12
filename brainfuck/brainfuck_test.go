package main

import (
	"math"
	"testing"
)

func TestAbs(t *testing.T) {
	// Default test, doesn't actually have anything to do with brainfuck
	expected := 1
	actual := int(math.Abs(-1))
	if expected != actual {
		t.Errorf("Abs(-1) = %d; want 1", actual)
	}
}

func TestGetBFCode(t *testing.T) {
	// Test file io
	var expected string = "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."
	var actual string = getBFCode("../Examples/helloWorld.bf")

	if expected != actual {
		t.Errorf("Expected: %s\n Actual: %s", expected, actual)
	}
}
