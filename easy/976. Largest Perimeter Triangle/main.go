package main

import (
	"fmt"
	"sort"
)

func largestPerimeter(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	sort.Ints(nums)
	right := len(nums) - 1

	for right >= 2 {
		if nums[right-2] > nums[right]-nums[right-1] {
			return nums[right-2] + nums[right-1] + nums[right]
		}

		right--
	}

	return 0
}

func main() {
	fmt.Println(largestPerimeter([]int{2, 1, 2}))
	fmt.Println(largestPerimeter([]int{1, 2, 1, 10}))
}
