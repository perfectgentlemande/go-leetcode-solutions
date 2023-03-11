package main

func findTargetSumWays(nums []int, target int) int {
	ways := 0

	var add func(sum int, i int)
	add = func(sum int, i int) {
		sum1 := sum - nums[i]
		sum2 := sum + nums[i]
		if i == len(nums)-1 {
			if sum1 == target {
				ways++
			}
			if sum2 == target {
				ways++
			}
			return
		}
		add(sum1, i+1)
		add(sum2, i+1)
	}

	add(0, 0)
	return ways
}

func main() {

}
