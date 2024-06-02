package main

func minimumChairs(s string) int {
	if s == "" {
		return 0
	}

	curCap := 0
	curMax := 0
	for _, ch := range s {
		if ch == 'E' {
			curCap++
		}

		if ch == 'L' {
			curCap--
		}

		if curCap > curMax {
			curMax = curCap
		}
	}

	return curMax
}

func main() {

}
