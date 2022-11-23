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

func removeDuplicatesBetter(nums []int) int {
	count := 0
	for i := range nums {
		if i < len(nums)-1 && nums[i] == nums[i+1] {
			continue
		}

		nums[count] = nums[i]
		count++
	}

	return count
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

	arr1 = []int{1, 1, 2}
	fmt.Println(arr1)
	fmt.Println(removeDuplicatesBetter(arr1))
	fmt.Println(arr1)

	arr2 = []int{0, 0, 1, 1, 2, 3, 4}
	fmt.Println(arr2)
	fmt.Println(removeDuplicatesBetter(arr2))
	fmt.Println(arr2)

	arr3 = []int{0, 1, 1, 1, 1, 2, 2}
	fmt.Println(arr3)
	fmt.Println(removeDuplicatesBetter(arr3))
	fmt.Println(arr3)
}
