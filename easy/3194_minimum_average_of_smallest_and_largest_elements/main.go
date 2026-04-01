package main

func minimumAverage(nums []int) float64 {
	var averages []float64
	minAverage := 50.5

	for len(nums) > 0 {
		minI := 0
		maxI := 0

		for i := range nums {
			if nums[i] < nums[minI] {
				minI = i
			}

			if nums[i] > nums[maxI] {
				maxI = i
			}
		}

		average := (float64(nums[minI]) + float64(nums[maxI])) / 2
		if len(averages) == 0 {
			minAverage = average
		}

		if average < minAverage {
			minAverage = average
		}
		averages = append(averages, average)

		newNums := make([]int, 0, len(nums))
		for i := range nums {
			if i != maxI && i != minI {
				newNums = append(newNums, nums[i])
			}
		}

		nums = newNums
	}

	return minAverage
}

func main() {

}
