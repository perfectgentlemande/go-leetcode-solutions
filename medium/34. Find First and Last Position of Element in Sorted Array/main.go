package main

import "fmt"

func findFirst(nums []int, target int) int {
	l, r := 0, len(nums)-1
	first := -1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			first = mid
			r = mid - 1
		} else if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return first
}
func findLast(nums []int, target int) int {
	l, r := 0, len(nums)-1
	last := -1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			last = mid
			l = mid + 1
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return last
}
func searchRange(nums []int, target int) []int {
	return []int{findFirst(nums, target), findLast(nums, target)}
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
}
