package main

type Cell struct {
	i, j int
}
type CellWithDistance struct {
	Cell
	distance int
}
type QueueNode struct {
	Val  CellWithDistance
	Next *QueueNode
}
type Queue struct {
	root *QueueNode
	rear *QueueNode
	size int
}

func (q *Queue) Add(val CellWithDistance) {
	nodePtr := &QueueNode{
		Val: val,
	}

	if q.root == nil {
		q.root = nodePtr
	}

	if q.rear != nil {
		q.rear.Next = nodePtr
	}

	q.rear = nodePtr
	q.size++
}

func (q *Queue) Remove() (CellWithDistance, bool) {
	if q.root == nil {
		return CellWithDistance{}, false
	}

	curVal := q.root.Val
	if q.root == q.rear {
		q.rear = nil
	}
	q.root = q.root.Next
	q.size--

	return curVal, true
}

func (q *Queue) Head() (CellWithDistance, bool) {
	if q.root == nil {
		return CellWithDistance{}, false
	}

	return q.root.Val, true
}

func (q *Queue) IsEmpty() bool {
	return q.root == nil
}

func (q *Queue) Size() int {
	return q.size
}

func updateMatrix(mat [][]int) [][]int {
	height := len(mat)
	width := len(mat[0])

	seen := make(map[Cell]bool, height)
	queue := Queue{}

	for i := range mat {
		for j := range mat[i] {
			if mat[i][j] == 0 {
				seen[Cell{i: i, j: j}] = true
				queue.Add(CellWithDistance{Cell: Cell{i: i, j: j}, distance: 1})
			}
		}
	}

	for !queue.IsEmpty() {
		cur, _ := queue.Remove()

		for _, d := range []Cell{{i: 0, j: 1}, {i: 1, j: 0}, {i: 0, j: -1}, {i: -1, j: 0}} {
			newCell := Cell{i: cur.i + d.i, j: cur.j + d.j}

			if newCell.i < 0 ||
				newCell.i >= height ||
				newCell.j < 0 ||
				newCell.j >= width ||
				seen[Cell{i: newCell.i, j: newCell.j}] {
				continue
			}

			seen[Cell{i: newCell.i, j: newCell.j}] = true
			queue.Add(CellWithDistance{Cell: Cell{i: newCell.i, j: newCell.j}, distance: cur.distance + 1})
			mat[newCell.i][newCell.j] = cur.distance
		}
	}

	return mat
}

func main() {
}
