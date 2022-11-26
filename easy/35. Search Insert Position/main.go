package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := (left + right) / 2

	for left <= right {
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}

		mid = (left + right) / 2
	}

	return left
}

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 1))
	fmt.Println(searchInsert(nums, 3))
	fmt.Println(searchInsert(nums, 5))
	fmt.Println(searchInsert(nums, 6))
	fmt.Println(searchInsert(nums, 2))
	fmt.Println(searchInsert(nums, 7))
}
