// https://leetcode.com/problems/unique-morse-code-words/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	arr := []string{"gin", "zen", "gun", "men"}
	uniqueMorseRepresentations(arr)
}

func uniqueMorseRepresentations(words []string) int {

	if len(words) < 1 || len(words) > 100 {
		return 0
	}

	for i := 0; i < len(words); i++ {
		if len(words[i]) > 1 && len(words[i]) < 12 {
			for j := 0; j < len(words[i]); j++ {
				if unicode.IsLetter(rune(words[i][j])) != true {
					return 0
				}
			}
		} else {
			return 0
		}
	}

	codes := [26]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	for i := 0; i < len(words); i++ {
		codeArray := []string{}

		for j := 0; j < len(words[i]); j++ {
			fmt.Println(string(words[i][j]))
			codeArray = append(codeArray, codes[words[i][j]-97])
		}
		fmt.Println(len(codeArray))
	}

	return 0
}
