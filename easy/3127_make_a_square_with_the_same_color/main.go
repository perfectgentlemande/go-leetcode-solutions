package main

func topLeftExists(grid [][]byte, i, j int, color byte) bool {
	if i == 0 {
		return false
	}

	if j == 0 {
		return false
	}

	return grid[i][j-1] == color && grid[i-1][j] == color && grid[i-1][j-1] == color
}

func topRightExists(grid [][]byte, i, j int, color byte) bool {
	if i == 0 {
		return false
	}

	if j == len(grid[i])-1 {
		return false
	}

	return grid[i][j+1] == color && grid[i-1][j] == color && grid[i-1][j+1] == color
}

func bottomLeftExists(grid [][]byte, i, j int, color byte) bool {
	if i == len(grid)-1 {
		return false
	}

	if j == 0 {
		return false
	}

	return grid[i][j-1] == color && grid[i+1][j] == color && grid[i+1][j-1] == color
}

func bottomRightExists(grid [][]byte, i, j int, color byte) bool {
	if i == len(grid)-1 {
		return false
	}

	if j == len(grid[i])-1 {
		return false
	}

	return grid[i][j+1] == color && grid[i+1][j] == color && grid[i+1][j+1] == color
}

func canMakeSquare(grid [][]byte) bool {
	if len(grid) == 1 {
		return false
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'B' || grid[i][j] == 'W' {
				if topLeftExists(grid, i, j, 'W') || topRightExists(grid, i, j, 'W') ||
					bottomLeftExists(grid, i, j, 'W') || bottomRightExists(grid, i, j, 'W') {
					return true
				}
			}

			if grid[i][j] == 'W' || grid[i][j] == 'B' {
				if topLeftExists(grid, i, j, 'B') || topRightExists(grid, i, j, 'B') ||
					bottomLeftExists(grid, i, j, 'B') || bottomRightExists(grid, i, j, 'B') {
					return true
				}
			}
		}
	}
	return false
}

func count(grid [][]byte, i, j int) (int, int) {
	if grid[i][j] == 'W' {
		return 0, 1
	}

	return 1, 0
}

func canMakeSquareBetter(grid [][]byte) bool {
	if len(grid) == 1 {
		return false
	}

	for i := 0; i < len(grid)-1; i++ {
		for j := 0; j < len(grid[i])-1; j++ {
			c1, c2 := count(grid, i, j)
			c11, c22 := count(grid, i+1, j)
			c111, c222 := count(grid, i, j+1)
			c1111, c2222 := count(grid, i+1, j+1)

			if c1+c11+c111+c1111 <= 1 ||
				c2+c22+c222+c2222 <= 1 {
				return true
			}
		}
	}

	return false
}

func main() {

}
