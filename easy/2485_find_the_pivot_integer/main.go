package main

import "fmt"

func pivotInteger(n int) int {
	left, right := 1, n
	leftSum, rightSum := left, right

	for left != right {
		if leftSum < rightSum {
			left++
			leftSum += left
		} else {
			right--
			rightSum += right
		}
	}

	if rightSum == leftSum {
		return left
	}

	return -1
}

func main() {
	fmt.Println(pivotInteger(8))
	fmt.Println(pivotInteger(1))
	fmt.Println(pivotInteger(4))
}
