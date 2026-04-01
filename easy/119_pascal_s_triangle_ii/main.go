package main

func getRow(rowIndex int) []int {
	res := make([][]int, 0, rowIndex+1)

	for i := 0; i <= rowIndex; i++ {
		res = append(res, make([]int, i+1))

		for j := 0; j < i+1; j++ {
			if j == 0 || i == j {
				res[i][j] = 1
			} else {
				res[i][j] = res[i-1][j-1] + res[i-1][j]
			}
		}
	}

	return res[rowIndex]
}

func getRowBetter(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	res[0] = 1

	for i := 1; i <= rowIndex; i++ {
		for j := i; j >= 1; j-- {
			res[j] += res[j-1]
		}
	}

	return res
}

func getNum(i, j int) int {
	if j == 0 || i == j {
		return 1
	}

	return getNum(i-1, j-1) + getNum(i-1, j)
}
func getRowRecursive(rowIndex int) []int {
	res := make([]int, 0, rowIndex+1)

	for i := 0; i <= rowIndex; i++ {
		res = append(res, getNum(rowIndex, i))
	}

	return res
}

func getRowRecursiveBetter(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	prevRow := getRowRecursiveBetter(rowIndex - 1)
	row := make([]int, rowIndex+1)

	for i := 0; i <= rowIndex; i++ {
		var left, right int

		if i > 0 {
			left = prevRow[i-1]
		}

		if i < rowIndex {
			right = prevRow[i]
		}

		row[i] = left + right
	}

	return row
}

func main() {

}
