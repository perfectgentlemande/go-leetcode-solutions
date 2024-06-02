package main

func minimumChairs(s string) int {
	res := 0

	if s == "" {
		return res
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
