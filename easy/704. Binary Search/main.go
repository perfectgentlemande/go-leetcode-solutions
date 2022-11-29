package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
	}

	return -1
}

func main() {
	fmt.Println([]int{-1, 0, 3, 5, 9, 12}, 9)
	fmt.Println([]int{-1, 0, 3, 5, 9, 12}, 2)
}
