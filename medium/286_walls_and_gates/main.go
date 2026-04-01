package main

import "math"

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

func wallsAndGates(rooms [][]int) {
	queue := Queue{}

	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[i]); j++ {
			if rooms[i][j] == 0 {
				queue.Add(Cell{i, j})
			}
		}
	}

	for !queue.IsEmpty() {
		size := len(queue)
		for k := 0; k < size; k++ {
			cell, _ := queue.Remove()

			i := cell.i
			j := cell.j

			// down
			if i < len(rooms)-1 && rooms[i+1][j] == math.MaxInt32 {
				queue.Add(Cell{i + 1, j})
				rooms[i+1][j] = rooms[i][j] + 1
			}

			// right
			if j < len(rooms[i])-1 && rooms[i][j+1] == math.MaxInt32 {
				queue.Add(Cell{i, j + 1})
				rooms[i][j+1] = rooms[i][j] + 1
			}

			// up
			if i > 0 && rooms[i-1][j] == math.MaxInt32 {
				queue.Add(Cell{i - 1, j})
				rooms[i-1][j] = rooms[i][j] + 1
			}

			// left
			if j > 0 && rooms[i][j-1] == math.MaxInt32 {
				queue.Add(Cell{i, j - 1})
				rooms[i][j-1] = rooms[i][j] + 1
			}
		}
	}
}
func main() {

}
