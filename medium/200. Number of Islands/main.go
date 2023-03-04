package main

var directions = [][]int{
	{1, 0},
	{-1, 0},
	{0, -1},
	{0, 1},
}

func isValid(y int, x int, grid [][]byte, visited [][]bool) bool {
	if y >= len(grid) || y < 0 {
		return false
	} else if x >= len(grid[0]) || x < 0 {
		return false
	} else if grid[y][x] == '0' {
		return false
	} else if visited[y][x] {
		return false
	}

	return true
}

type Cell struct {
	i, j int
}

type Queue []Cell

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}
func (q *Queue) Add(cell Cell) {
	*q = append([]Cell{cell}, *q...)
}
func (q *Queue) Remove() (Cell, bool) {
	if q.IsEmpty() {
		return Cell{}, false
	} else {
		index := len(*q) - 1
		element := (*q)[index]
		*q = (*q)[:index]
		return element, true
	}
}

func numIslands(grid [][]byte) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}

	numIslands := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				queue := Queue{}
				queue.Add(Cell{i, j})

				for !queue.IsEmpty() {
					curr, _ := queue.Remove()
					for _, direction := range directions {
						x := curr.i + direction[0]
						y := curr.j + direction[1]

						if isValid(x, y, grid, visited) {
							visited[x][y] = true
							queue = append(queue, Cell{x, y})
						}
					}
				}
				numIslands++
			}
		}
	}
	return numIslands
}
func main() {

}
