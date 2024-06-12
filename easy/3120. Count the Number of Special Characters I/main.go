package main

import (
	"unicode"
)

func numberOfSpecialChars(word string) int {
	mappedCharsLower := map[rune]struct{}{}
	mappedCharsNotLower := map[rune]struct{}{}
	count := 0

	for _, ch := range word {
		if unicode.IsLower(ch) {
			mappedCharsLower[ch] = struct{}{}
		} else {
			mappedCharsNotLower[unicode.ToLower(ch)] = struct{}{}
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
