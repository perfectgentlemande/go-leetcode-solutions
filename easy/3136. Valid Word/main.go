package main

import "unicode"

func isValid(word string) bool {
	vowels := map[rune]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}

	foundConsonant := false
	foundVowel := false
	for _, ch := range word {
		if !unicode.IsDigit(ch) && !unicode.IsLetter(ch) {
			return false
		}

		if _, ok := vowels[ch]; ok {
			foundVowel = true
		} else {
			foundConsonant = true
		}
	}

	return foundConsonant == foundVowel
}

func main() {

}
