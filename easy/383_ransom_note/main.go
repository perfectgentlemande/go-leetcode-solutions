package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	mappedWord := map[rune]int{}

	for _, note := range ransomNote {
		mappedWord[note]++
	}

	for _, ch := range magazine {
		if v, ok := mappedWord[ch]; ok && v != 0 {
			mappedWord[ch]--
		}
	}

	for _, count := range mappedWord {
		if count != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
}
