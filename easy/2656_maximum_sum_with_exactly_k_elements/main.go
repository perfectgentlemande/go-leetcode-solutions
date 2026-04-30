package main

import (
	"sort"
)

func countSum(nums []int) int {
	res := 0

	for _, v := range nums {
		res += v
	}

	return res
}
func maximizeSum(nums []int, k int) int {
	sort.Ints(nums)
	resArr := make([]int, 0, k)

	for i := 0; i < k; i++ {
		resArr = append(resArr, nums[len(nums)-1])
		nums[len(nums)-1] = nums[len(nums)-1] + 1
	}

	return countSum(resArr)
}

func main() {}
