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

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reversedDigits := extractReversedDigits(x)
	for i := range reversedDigits {
		if reversedDigits[i] != reversedDigits[len(reversedDigits)-1-i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}
