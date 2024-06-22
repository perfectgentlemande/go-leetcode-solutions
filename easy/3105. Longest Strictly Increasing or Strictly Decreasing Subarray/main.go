package main

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func longestMonotonicSubarray(nums []int) int {
	curMaxIncLen := 1
	curMaxDecLen := 1

	for i := 0; i < len(nums); i++ {
		increasingLen := 1
		decreasingLen := 1

		prev := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[prev] {
				increasingLen++
			} else {
				curMaxIncLen = max(increasingLen, curMaxIncLen)
				increasingLen = 1
			}

			if nums[j] < nums[prev] {
				decreasingLen++
			} else {
				curMaxDecLen = max(decreasingLen, curMaxDecLen)
				decreasingLen = 1
			}

			prev++
		}

		curMaxIncLen = max(increasingLen, curMaxIncLen)
		curMaxDecLen = max(decreasingLen, curMaxDecLen)
	}

	return max(curMaxIncLen, curMaxDecLen)
}

func main() {

}
