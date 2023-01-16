package main

import (
	"fmt"
	"sort"
)

func areAlmostEqual(s1 string, s2 string) bool {
	counter := 0

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			counter++
		}

		if counter > 2 {
			return false
		}
	}

	s1Runes, s2Runes := []rune(s1), []rune(s2)
	sort.Slice(s1Runes, func(i int, j int) bool { return s1Runes[i] < s1Runes[j] })
	sort.Slice(s2Runes, func(i int, j int) bool { return s2Runes[i] < s2Runes[j] })

	if string(s1Runes) == string(s2Runes) {
		return true
	}

	return counter == 0
}

func main() {
	fmt.Println("bank", "kanb")
	fmt.Println("attack", "defend")
	fmt.Println("kelb", "kelb")
}
