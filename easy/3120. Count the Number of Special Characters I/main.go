package main

import (
	"unicode"
)

func numberOfSpecialChars(word string) int {
	mappedCharsLower := map[rune]int{}
	mappedCharsNotLower := map[rune]int{}
	count := 0

	for _, ch := range word {
		if unicode.IsLower(ch) {
			mappedCharsLower[ch]++
		} else {
			mappedCharsNotLower[unicode.ToLower(ch)]++
		}
	}

	for k, _ := range mappedCharsLower {
		if _, ok := mappedCharsNotLower[k]; ok {
			count++
		}
	}

	return count
}

func main() {

}
