package main

func findValidPair(s string) string {
	runes := []rune(s)

	freqs := make(map[rune]int)
	for _, n := range runes {
		freqs[n]++
	}

	for i := 0; i < len(runes)-1; i++ {
		if runes[i] == runes[i+1] {
			continue
		}

		if freqs[runes[i]] == int(s[i])-48 && freqs[runes[i+1]] == int(s[i+1])-48 {
			return string([]rune{runes[i], runes[i+1]})
		}
	}

	return ""
}

func main() {

}
