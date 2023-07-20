package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrimeFactors(t *testing.T) {
	testCases := []struct {
		input    int
		expected []int
	}{
		{input: 1, expected: []int{}},
		{input: 2, expected: []int{2}},
		{input: 3, expected: []int{3}},
		{input: 4, expected: []int{2, 2}},
		{input: 12, expected: []int{2, 2, 3}},
		{input: 84, expected: []int{2, 2, 3, 7}},
		{input: 100, expected: []int{2, 2, 5, 5}},
		{input: 101, expected: []int{101}},
		{input: 1024, expected: []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %d", tc.input), func(t *testing.T) {
			actual := PrimeFactors(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("PrimeFactors(%d) = %v, expected %v", tc.input, actual, tc.expected)
			}
		})
	}
}
