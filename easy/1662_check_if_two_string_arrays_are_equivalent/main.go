package main

import "strings"

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	builder1 := strings.Builder{}
	builder2 := strings.Builder{}

	for _, word := range word1 {
		builder1.WriteString(word)
	}
	for _, word := range word2 {
		builder2.WriteString(word)
	}

	return builder1.String() == builder2.String()
}

func main() {

}
