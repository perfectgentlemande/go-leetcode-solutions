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

func specialArrayBinary(nums []int) int {
	res := -1

	max := findMax(nums)
	l, r := 0, max
	for l <= r {
		mid := l + (r-l)/2

		counter := 0
		for j := range nums {
			if nums[j] >= mid {
				counter++
			}
		}

		if counter == mid {
			return counter
		}

		if counter > mid {
			l = mid + 1
		}

		if counter < mid {
			r = mid - 1
		}
	}

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
	fmt.Println(specialArrayBinary([]int{3, 5}))
	fmt.Println(specialArrayBinary([]int{0, 0}))
	fmt.Println(specialArrayBinary([]int{0, 4, 3, 0, 4}))
}
