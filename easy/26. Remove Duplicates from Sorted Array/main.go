package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	leftBoundary := len(nums) - 1
	for i := 1; i <= leftBoundary; i++ {
		for nums[i-1] == nums[i] {
			for j := i; j < len(nums)-1; j++ {
				nums[j-1] = nums[j]
				nums[j] = nums[leftBoundary]
			}
			nums[leftBoundary] = -1000
			leftBoundary--
		}
	}

	return leftBoundary + 1
}

func main() {
	arr1 := []int{1, 1, 2}
	fmt.Println(arr1)
	fmt.Println(removeDuplicates(arr1))
	fmt.Println(arr1)

	arr2 := []int{0, 0, 1, 1, 2, 3, 4}
	fmt.Println(arr2)
	fmt.Println(removeDuplicates(arr2))
	fmt.Println(arr2)

	arr3 := []int{0, 1, 1, 1, 1, 2, 2}
	fmt.Println(arr3)
	fmt.Println(removeDuplicates(arr3))
	fmt.Println(arr3)
}
