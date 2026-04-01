package main

import (
	"fmt"
	"sort"
)

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)

	if len(arr) <= 2 {
		return true
	}

	for i := len(arr) - 1; i >= 2; i-- {
		if arr[i]-arr[i-1] != arr[i-1]-arr[i-2] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(canMakeArithmeticProgression([]int{3, 5, 1}))
	fmt.Println(canMakeArithmeticProgression([]int{1, 2, 4}))
}
