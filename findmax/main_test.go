package main

import "testing"

func TestFindMax(t *testing.T) {
	// Test case with positive integers
	arr := []int{4, 7, 1, 9, 2}
	expectedMax := 9
	if max := FindMax(arr); max != expectedMax {
		t.Errorf("Expected max %d, but got %d", expectedMax, max)
	}

	// Test case with negative integers
	arr = []int{-2, -9, -3, -1, -7}
	expectedMax = -1
	if max := FindMax(arr); max != expectedMax {
		t.Errorf("Expected max %d, but got %d", expectedMax, max)
	}

	// Test case with mixed positive and negative integers
	arr = []int{-5, 2, -8, 10, 3}
	expectedMax = 10
	if max := FindMax(arr); max != expectedMax {
		t.Errorf("Expected max %d, but got %d", expectedMax, max)
	}

	// Test case with an empty array
	arr = []int{}
	expectedMax = 0 // As we are considering max as 0 for an empty array
	if max := FindMax(arr); max != expectedMax {
		t.Errorf("Expected max %d, but got %d", expectedMax, max)
	}
}
