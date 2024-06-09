package main

import "unicode"

unc isValid(word string) bool {
    if len(word) < 3 {
        return false
    }

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

        if unicode.IsLetter(ch) {
            if _, ok := vowels[ch]; ok {
			foundVowel = true
            } else {
                foundConsonant = true
            }
        }	
	}

	return foundConsonant && foundVowel
}

func main() {

}
