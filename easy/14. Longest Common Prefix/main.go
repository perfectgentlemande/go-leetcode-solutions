package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	resRunes := make([]rune, 0)
	for chIndex, ch := range strs[0] {
		willBeAdded := true

		for i := 1; i < len(strs); i++ {
			if chIndex >= len(strs[i]) || ch != rune(strs[i][chIndex]) {
				willBeAdded = false
				break
			}
		}

		if willBeAdded {
			resRunes = append(resRunes, ch)
		} else {
			break
		}
	}

	return string(resRunes)
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
