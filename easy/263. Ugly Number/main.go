package main

import "fmt"

func maxDivide(a, b int) int {
	for a%b == 0 {
		a = a / b
	}

	return a
}
func isUgly(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}

	for _, pf := range []int{2, 3, 5} {
		for {
			if n%pf != 0 {
				break
			}
			n = n / pf
		}
	}

	return n == 1
}

func main() {
	fmt.Println(isUgly(6))
	fmt.Println(isUgly(1))
	fmt.Println(isUgly(14))
}
