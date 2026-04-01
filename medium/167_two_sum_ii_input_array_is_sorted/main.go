package main

import "fmt"

func binarySearch(numbers []int, target int) int {
	l, r := 0, len(numbers)-1
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
func twoSum(numbers []int, target int) []int {
	res := []int{-1, -1}

	for i := range numbers {
		if numbers[i]+numbers[i] == target && i+1 < len(numbers) && numbers[i+1] == numbers[i] {
			res[0], res[1] = i+1, i+2
			return res
		}

		another := binarySearch(numbers, target-numbers[i])
		if another >= 0 {
			res[0], res[1] = i+1, another+1
			return res
		}
	}

	return res
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
	fmt.Println(twoSum([]int{-1, 0}, -1))
}
