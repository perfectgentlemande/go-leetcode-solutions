package main

import (
	"fmt"
)

func signFunc(x int) int {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}

	return 0
}

func arraySign(nums []int) int {
	res := 1

	for _, num := range nums {
		res *= num
	}

	return signFunc(res)
}

func main() {
	fmt.Println(arraySign([]int{-1, -2, -3, -4, 3, 2, 1}))
	fmt.Println(arraySign([]int{1, 5, 0, 2, -3}))
	fmt.Println(arraySign([]int{-1, 1, -1, 1, -1}))
}
