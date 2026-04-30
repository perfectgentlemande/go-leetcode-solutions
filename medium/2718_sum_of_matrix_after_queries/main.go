package main

func setColumnToValue(matrix [][]int, column int, value int) {
	for i := range matrix {
		matrix[i][column] = value
	}
}

func setRowToValue(matrix [][]int, row int, value int) {
	for j := range matrix[row] {
		matrix[row][j] = value
	}
}

func countSum(matrix [][]int) int64 {
	var result int64

	for i := range matrix {
		for j := range matrix[i] {
			result += int64(matrix[i][j])
		}
	}

	return result
}

func matrixSumQueriesBrute(n int, queries [][]int) int64 {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for i := range queries {
		if queries[i][0] == 0 {
			setRowToValue(matrix, queries[i][1], queries[i][2])
		}
		if queries[i][0] == 1 {
			setColumnToValue(matrix, queries[i][1], queries[i][2])
		}
	}

	return countSum(matrix)
}

func matrixSumQueries(n int, queries [][]int) int64 {
	sum := int64(0)
	visitedRows := make(map[int]bool)
	visitedColumns := make(map[int]bool)
	remainingRows := n
	remainingColumns := n

	for i := len(queries) - 1; i >= 0; i-- {
		if queries[i][0] == 0 {
			if visitedRows[queries[i][1]] {
				continue
			}

			sum += int64(remainingColumns) * int64(queries[i][2])
			visitedRows[queries[i][1]] = true
			remainingRows--
		}

		if queries[i][0] == 1 {
			if visitedColumns[queries[i][1]] {
				continue
			}

			sum += int64(remainingRows) * int64(queries[i][2])
			visitedColumns[queries[i][1]] = true
			remainingColumns--
		}
	}

	return sum
}

func main() {

}
