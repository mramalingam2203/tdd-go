// Write a function that takes a string as input and returns the reverse of that string.

package main

import "fmt"

func main() {
	s := "hello world"
	fmt.Println(stringReverse(s))
}

func stringReverse(s string) string {
	result := []byte{}
	for i := 0; i < len(s); i++ {
		result = append(result, s[len(s)-(i+1)])
	}
	return string(result)
}
