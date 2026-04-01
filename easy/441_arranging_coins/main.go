package main

import "fmt"

func arrangeCoins(n int) int {
	i := 1
	for n > 0 {
		// fmt.Printf("n: %d, i: %d\n", n, i)

		n = n - i

		if n < 0 {
			return i - 1
		}
		if n == 0 {
			return i
		}

		i++
	}

	return 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func arrangeCoinsBinary(n int) int {
	l, r := 1, n
	res := 0

	for l <= r {
		mid := l + (r-l)/2
		sum := (mid * (mid + 1)) / 2

		if sum > n {
			r = mid - 1
		} else {
			l = mid + 1
			res = max(res, mid)
		}
	}

	return res
}

func main() {
	fmt.Println(arrangeCoins(5))
	fmt.Println(arrangeCoins(8))
	fmt.Println(arrangeCoins(1))
	fmt.Println(arrangeCoins(2089018456))
	fmt.Println(arrangeCoinsBinary(5))
	fmt.Println(arrangeCoinsBinary(8))
	fmt.Println(arrangeCoinsBinary(1))
	fmt.Println(arrangeCoinsBinary(2089018456))
}
