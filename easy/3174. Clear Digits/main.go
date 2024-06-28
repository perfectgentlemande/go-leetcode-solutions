package main

import (
	"unicode"
)

func clearDigits(s string) string {
	if s == "" {
		return ""
	}

	prevRuneStr := []rune(s)
	newRuneStr := make([]rune, 0, len(s))
	last := false
	steps := 0

	for !last {
		if len(prevRuneStr) == 0 {
			return string(prevRuneStr)
		}

		for i, rn := range prevRuneStr {
			if unicode.IsDigit(rn) {
				if i != 0 {
					if i != len(prevRuneStr)-1 {
						newRuneStr = append(prevRuneStr[:i-1], prevRuneStr[i+1:]...)
					} else {
						newRuneStr = prevRuneStr[:i-1]
					}
				}

				if i == len(prevRuneStr) {
					last = true
				}

				break
			} else {
				if i == len(prevRuneStr)-1 {
					return string(prevRuneStr)
				}
			}
		}

		prevRuneStr = newRuneStr
		steps++
	}

	return string(prevRuneStr)
}

func clearDigitsBetter(s string) string {
	var stack []rune

	for _, rn := range s {
		if rn >= 'a' && rn <= 'z' {
			stack = append(stack, rn)
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return string(stack)
}

func main() {

}
