/*
	Create a function that takes an integer as input and returns its prime factors in an array.
	For example, if the input is 12, the function should return [2, 2, 3], as 2 * 2 * 3 = 12.
*/

package main

import "fmt"

func PrimeFactors(n int) []int {
	factors := []int{}

	// Check for 2 as a factor and divide until n is odd
	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}

	// Check for other odd prime factors
	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}

	// If n is a prime number greater than 2, add it as a factor
	if n > 2 {
		factors = append(factors, n)
	}

	return factors
}

func main() {
	number := 84
	factors := PrimeFactors(number)
	fmt.Printf("Prime factors of %d: %v\n", number, factors)
}
