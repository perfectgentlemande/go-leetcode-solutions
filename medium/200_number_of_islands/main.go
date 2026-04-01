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

type Node struct {
	Val  []int
	Next *Node
}
type Stack struct {
	root *Node
}

func (s *Stack) Push(val []int) {
	nextRoot := s.root

	s.root = &Node{
		Val:  val,
		Next: nextRoot,
	}
}

func (s *Stack) Pop() {
	if s.root == nil {
		return
	}

	s.root = s.root.Next
}

func (s *Stack) Top() []int {
	if s.root == nil {
		return nil
	}

	return s.root.Val
}

func (s *Stack) IsEmpty() bool {
	return s.root == nil
}

const (
	// LAND ...
	LAND = 49
	// WATER ...
	WATER = 48
)

func numIslandsDFS(grid [][]byte) int {
	visited := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]int, len(grid[i]))
	}

	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if (grid[i][j] == LAND) && (visited[i][j] == 0) {
				// explore on it, and mark the current island as visited
				// so we don't re-count it
				markIsland(grid, visited, []int{i, j})
				count++
			}
		}
	}

	return count
}

// pos x,y
func markIsland(grid [][]byte, visited [][]int, pos []int) {
	height := len(grid)
	width := len(grid[0])

	explore := Stack{}
	explore.Push(pos)

	for !explore.IsEmpty() {
		// pop the last element in the stack
		position := explore.Top()
		explore.Pop()

		i, j := position[0], position[1]

		// top
		if i >= 1 && i-1 <= height-1 && grid[i-1][j] == LAND && visited[i-1][j] == 0 {
			explore.Push([]int{i - 1, j})
		}

		// left
		if j >= 1 && j-1 <= width-1 && grid[i][j-1] == LAND && visited[i][j-1] == 0 {
			explore.Push([]int{i, j - 1})
		}

		// right
		if j+1 <= width-1 && grid[i][j+1] == LAND && visited[i][j+1] == 0 {
			explore.Push([]int{i, j + 1})
		}

		// bottom
		if i+1 <= height-1 && grid[i+1][j] == LAND && visited[i+1][j] == 0 {
			explore.Push([]int{i + 1, j})
		}

		// mark the current element as explored
		visited[i][j] = 1
	}
}

func main() {

}
