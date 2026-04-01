package main

var climbResults map[int]int = map[int]int{}

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	res, ok := climbResults[n]
	if !ok {
		res = climbStairs(n-1) + climbStairs(n-2)
		climbResults[n] = res
	}

	return res
}

func climbStairsIterative(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	previous := 2
	previousPrevious := 1
	current := 0

	for i := 3; i <= n; i++ {
		current = previous + previousPrevious
		previousPrevious = previous
		previous = current
	}

	return current
}
func main() {

}
