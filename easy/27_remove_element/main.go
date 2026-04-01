package main

import "fmt"

func removeElement(nums []int, val int) int {
	count := 0

	i := 0
	for i < len(nums) {
		if nums[i] == val {
			count++
		} else {
			nums[i-count] = nums[i]
		}
		i++
	}

	return i - count
}
func main() {
	arr1 := []int{3, 2, 2, 3}
	fmt.Println(arr1)
	fmt.Println(removeElement(arr1, 3))
	fmt.Println(arr1)

	arr2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(arr2)
	fmt.Println(removeElement(arr2, 2))
	fmt.Println(arr2)
}
