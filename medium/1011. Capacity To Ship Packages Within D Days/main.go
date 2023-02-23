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

func shipWithinDays(weights []int, days int) int {
	cap := findMax(weights)
	if len(weights) == 0 {
		return 0
	}

	for {
		curDays := 1
		curPack := 0

		for _, w := range weights {
			curPack += w

			if curPack > cap {
				curPack = w
				curDays++
			}
		}

		if curDays <= days {
			return cap
		}

		cap++
	}
}

func main() {

}
