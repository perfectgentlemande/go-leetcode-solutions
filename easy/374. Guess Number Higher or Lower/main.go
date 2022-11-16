package main

import "fmt"

var pick int

func guess(num int) int {
	if num > pick {
		return -1
	}

	if num < pick {
		return 1
	}

	return 0
}
func guessNumber(n int) int {
	left, right := 1, n

	for left <= n {
		mid := left + (right-left)/2
		fmt.Println(mid)

		res := guess(mid)
		if res == 0 {
			return mid
		}

		if res < 0 {
			right = mid - 1
			continue
		}

		left = mid + 1
	}

	return 0
}

func main() {
	pick = 6
	fmt.Println("res:")
	fmt.Println(guessNumber(10))
	pick = 1
	fmt.Println("res:")
	fmt.Println(guessNumber(1))
	pick = 1
	fmt.Println("res:")
	fmt.Println(guessNumber(2))
}
