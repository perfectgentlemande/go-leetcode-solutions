package main

func hasKFreqs(symbols map[string]int, k int) bool {
	for _, count := range symbols {
		if count >= k {
			return true
		}
	}
	return false
}

func numberOfSubstrings(s string, k int) int {
	count := 0
	countMap := map[string]int{}
	start, end := 0, 0

	for end < len(s) {
		countMap[string(s[end])]++

		if hasKFreqs(countMap, k) {
			tempres := len(s) - end
			count += tempres
			track := 0

			for hasKFreqs(countMap, k) {
				track++
				countMap[string(s[start])]--
				start++
			}

			if track-1 >= 0 {
				count = count + ((track - 1) * (tempres))
			}
		}

		end++
	}

	return count
}

func main() {

}
