package main

import "fmt"

func isValid(s string) bool {
	bracketStack := make([]rune, 0, len(s))

	for _, ch := range s {
		length := len(bracketStack)

		switch {
		case ch == '(' || ch == '[' || ch == '{':
			bracketStack = append(bracketStack, ch)
		case ch == ')' && length != 0 && bracketStack[length-1] == '(':
			bracketStack = bracketStack[:length-1]
		case ch == '}' && length != 0 && bracketStack[length-1] == '{':
			bracketStack = bracketStack[:length-1]
		case ch == ']' && length != 0 && bracketStack[length-1] == '[':
			bracketStack = bracketStack[:length-1]
		default:
			return false
		}
	}

	return len(bracketStack) == 0
}
func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
}
