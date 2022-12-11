package main

import "fmt"

func findPivot(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)/2

		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}

func binarySearch(numbers []int, target, l, r int) int {
	for l <= r {
		mid := l + (r-l)/2

		if numbers[mid] == target {
			return mid
		}

		if numbers[mid] > target {
			r = mid - 1
		}

		if numbers[mid] < target {
			l = mid + 1
		}
	}

	return -1
}

func search(nums []int, target int) int {
	pivotIndex := findPivot(nums)

	if nums[pivotIndex] == target {
		return pivotIndex
	}

	if target < nums[pivotIndex] {
		return -1
	}

	if target <= nums[len(nums)-1] && pivotIndex < len(nums) {
		return binarySearch(nums, target, pivotIndex+1, len(nums)-1)
	}

	return binarySearch(nums, target, 0, pivotIndex-1)
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println(search([]int{1}, 0))
	fmt.Println(search([]int{1, 3}, 3))

	fmt.Println(findPivot([]int{1, 3}))
}
