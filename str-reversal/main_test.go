package main

import "testing"

func TestFindMax(t *testing.T) {
	// Test case with positive integers
	str := "hello world"
	expectedResult := "dlrow olleh"
	if resultant := stringReverse(str); expectedResult != resultant {
		t.Errorf("Expected reversal %s, but got %s", expectedResult, resultant)
	}

	str = "  hello world "
	expectedResult = " dlrow olleh  "

	if resultant := stringReverse(str); expectedResult != resultant {
		t.Errorf("Expected reversal %s, but got %s", expectedResult, resultant)
	}

}
