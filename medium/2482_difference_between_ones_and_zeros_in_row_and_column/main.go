package main

func newDiffGrid(grid [][]int) [][]int {
	diffGrid := make([][]int, len(grid))
	for i := range grid {
		diffGrid[i] = make([]int, len(grid[i]))
	}

	return diffGrid
}

func onesRow(grid [][]int, rowIndex int) int {
	count := 0

	for j := range grid[rowIndex] {
		if grid[rowIndex][j] == 1 {
			count++
		}
	}

	return count
}
func onesCol(grid [][]int, colIndex int) int {
	count := 0

	for i := range grid {
		if grid[i][colIndex] == 1 {
			count++
		}
	}

	return count
}
func zerosRow(grid [][]int, rowIndex int) int {
	count := 0

	for j := range grid[rowIndex] {
		if grid[rowIndex][j] == 0 {
			count++
		}
	}

	return count
}
func zerosCol(grid [][]int, colIndex int) int {
	count := 0

	for i := range grid {
		if grid[i][colIndex] == 0 {
			count++
		}
	}

	return count
}

func onesMinusZeros(grid [][]int) [][]int {
	diffGrid := newDiffGrid(grid)

	for i := range grid {
		for j := range grid[i] {
			diffGrid[i][j] = onesRow(grid, i) + onesCol(grid, j) - zerosRow(grid, i) - zerosCol(grid, j)
		}
	}

	return diffGrid
}

func onesMinusZerosBetter(grid [][]int) [][]int {
	n := len(grid)
	m := len(grid[0])
	col := make([]int, m)
	row := make([]int, n)

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				row[i]++
				col[j]++
			} else {
				row[i]--
				col[j]--
			}
		}
	}

	res := make([][]int, n)
	for i := range grid {
		res[i] = make([]int, m)

		for j := range grid[i] {
			res[i][j] = col[j] + row[i]
		}
	}

	return res
}

func main() {

}
