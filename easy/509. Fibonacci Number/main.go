package main

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

var fibResults map[int]int = map[int]int{}

func fibWithMemoization(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	res, ok := fibResults[n]
	if !ok {
		res = fibWithMemoization(n-1) + fibWithMemoization(n-2)
		fibResults[n] = res
	}

	return res
}

func main() {

}
