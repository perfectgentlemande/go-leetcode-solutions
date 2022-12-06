package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	row := 0
	col := len(matrix[row]) - 1

	for row < len(matrix) && col >= 0 {
		if matrix[row][col] == target {
			return true
		}

		if matrix[row][col] < target {
			row++
		} else if matrix[row][col] > target {
			col--
		}
	}

	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
	fmt.Println(searchMatrix([][]int{{1}}, 2))
}
