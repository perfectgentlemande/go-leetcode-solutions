package main

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

var directions = [][]int{
	{1, 0},
	{-1, 0},
	{0, -1},
	{0, 1},
}

func isEdge(y int, x int, grid [][]int) bool {
	if y >= len(grid) || y < 0 {
		return true
	}

	if x >= len(grid[y]) || x < 0 {
		return true
	}

	return false
}
func isValid(y int, x int, grid [][]int, visited [][]bool) bool {
	if isEdge(y, x, grid) {
		return false
	} else if grid[y][x] == 1 {
		return false
	} else if visited[y][x] {
		return false
	}

	return true
}

func closedIsland(grid [][]int) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}

	allIslands := 0
	numNotClosedIslands := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if !visited[i][j] && grid[i][j] == 0 {
				queue := Queue{}
				queue.Add(Cell{i, j})
				allIslands++
				notClosed := false

				for !queue.IsEmpty() {
					curr, _ := queue.Remove()

					for _, direction := range directions {
						x := curr.i + direction[0]
						y := curr.j + direction[1]

						if isEdge(x, y, grid) && !notClosed {
							numNotClosedIslands++
							notClosed = true
						}

						if isValid(x, y, grid, visited) {
							visited[x][y] = true
							queue.Add(Cell{x, y})
						}
					}
				}
			}
		}
	}

	return allIslands - numNotClosedIslands
}

func main() {

}
