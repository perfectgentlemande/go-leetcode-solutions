package main

import "fmt"

func extractReversedDigits(x int) []int {
	res := make([]int, 0)

	for x > 0 {
		res = append(res, x%10)
		x = x / 10
	}

	return res
}

func sumSquareDigits(ds []int) int {
	res := 0

	for _, num := range ds {
		res += num * num
	}

	return res
}

func isHappy(n int) bool {
	visited := map[int]struct{}{}

	for {
		sum := sumSquareDigits(extractReversedDigits(n))
		if sum == 1 {
			return true
		}
		if _, ok := visited[sum]; ok {
			return false
		}
		visited[sum] = struct{}{}

		n = sum
	}
}

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(2))
}
