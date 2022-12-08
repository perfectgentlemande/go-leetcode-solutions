package main

import "fmt"

func sqrt(x int) int {
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

func judgeSquareSum(c int) bool {
	l := 0
	r := sqrt(c)

	for l <= r {
		if l*l+r*r == c {
			return true
		}
		if l*l+r*r > c {
			r--
		}
		if l*l+r*r < c {
			l++
		}
	}

	return false
}

func main() {
	fmt.Println(judgeSquareSum(5))
	fmt.Println(judgeSquareSum(3))
}
