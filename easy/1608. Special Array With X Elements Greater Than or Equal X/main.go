package main

import "fmt"

func findMax(nums []int) int {
	max := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}
func specialArray(nums []int) int {
	res := -1

	max := findMax(nums)
	for i := 0; i <= max; i++ {
		counter := 0

		for j := range nums {
			if nums[j] >= i {
				counter++
			}
		}

		if counter == i {
			return counter
		}
	}

	return res
}

func main() {
	fmt.Println(specialArray([]int{3, 5}))
	fmt.Println(specialArray([]int{0, 0}))
	fmt.Println(specialArray([]int{0, 4, 3, 0, 4}))
}
