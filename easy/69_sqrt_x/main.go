package main

import "fmt"

func mySqrt(x int) int {
	l, r := 0, x

	for l <= r {
		mid := l + (r-l)/2
		if mid*mid == x {
			return mid
		}
		if (mid-1)*(mid-1) < x && mid*mid > x {
			return mid - 1
		}
		if mid*mid > x {
			r = mid - 1
		}
		if mid*mid < x {
			l = mid + 1
		}

	}
	return -1
}

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(7))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(10))
}
