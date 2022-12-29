package main

import "fmt"

func singleNumber(nums []int) int {
	res := 0

	for _, n := range nums {
		res ^= n
	}

	return res
}

func main() {
	fmt.Println([]int{2, 2, 1})
	fmt.Println([]int{4, 1, 2, 1, 2})
	fmt.Println([]int{1})
}
