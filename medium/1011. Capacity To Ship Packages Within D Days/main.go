package main

func findMax(weights []int) int {
	if len(weights) == 0 {
		return 0
	}

	res := weights[0]
	for _, w := range weights {
		if w > res {
			res = w
		}
	}

	return res
}

func checkCap(weights []int, cap, days int) bool {
	curDays := 1
	curPack := 0

	for _, w := range weights {
		curPack += w

		if curPack > cap {
			curPack = w
			curDays++
		}
	}

	return curDays <= days
}

func shipWithinDays(weights []int, days int) int {
	if len(weights) == 0 {
		return 0
	}

	cap := findMax(weights)

	for {
		if checkCap(weights, cap, days) {
			return cap
		}

		cap++
	}
}

func getMaxAndTotalLoad(weights []int) (max int, total int) {
	for _, w := range weights {
		if w > max {
			max = w
		}
		total += w
	}

	return
}

func shipWithinDaysBetter(weights []int, days int) int {
	if len(weights) == 0 {
		return 0
	}

	l, r := getMaxAndTotalLoad(weights)

	for l < r {
		mid := l + (r-l)/2

		if checkCap(weights, mid, days) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func main() {

}
