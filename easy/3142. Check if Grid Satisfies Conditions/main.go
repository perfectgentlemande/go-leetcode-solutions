package main

func satisfiesConditions(grid [][]int) bool {
	for i := range grid {
		for j := range grid[i] {
			if i < len(grid)-1 {
				if grid[i][j] != grid[i+1][j] {
					return false
				}
			}

			if j < len(grid[i])-1 {
				if grid[i][j] == grid[i][j+1] {
					return false
				}
			}
		}
	}

	return true
}

func main() {

}
