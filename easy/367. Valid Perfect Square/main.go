package main

import "fmt"

func isPerfectSquare(num int) bool {
	for i := 0; i*i <= num; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}
func isPerfectSquareBinary(num int) bool {
	l, r := 0, num

	for l <= r {
		mid := l + (r-l)/2
		if mid*mid == num {
			return true
		}
		if mid*mid > num {
			r = mid - 1
		}
		if mid*mid < num {
			l = mid + 1
		}
	}
	return false
}
func main() {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(14))
	fmt.Println(isPerfectSquareBinary(16))
	fmt.Println(isPerfectSquareBinary(14))
}
