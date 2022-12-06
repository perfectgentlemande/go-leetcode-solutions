package main

import "fmt"

func countNegatives(grid [][]int) int {
	count := 0

	for i := range grid {
		curCol := len(grid[i]) - 1

		for curCol >= 0 && grid[i][curCol] < 0 {
			curCol--
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(countNegatives([][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}))
	fmt.Println(countNegatives([][]int{{3, 2}, {1, 0}}))
}
