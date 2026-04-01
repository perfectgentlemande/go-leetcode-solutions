package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type QueueNode struct {
	Val  *TreeNode
	Next *QueueNode
}
type Queue struct {
	root *QueueNode
	rear *QueueNode
	size int
}

func (q *Queue) Add(val *TreeNode) {
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

func (q *Queue) Remove() (*TreeNode, bool) {
	if q.root == nil {
		return nil, false
	}

	curVal := q.root.Val
	if q.root == q.rear {
		q.rear = nil
	}
	q.root = q.root.Next
	q.size--

	return curVal, true
}

func (q *Queue) Head() (*TreeNode, bool) {
	if q.root == nil {
		return nil, false
	}

	return q.root.Val, true
}

func (q *Queue) IsEmpty() bool {
	return q.root == nil
}

func (q *Queue) Size() int {
	return q.size
}

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return false
	}

	q := Queue{}
	q.Add(root)
	end := false

	for !q.IsEmpty() {
		cur, _ := q.Remove()
		if cur == nil {
			end = true
		} else {
			if end {
				return false
			}

			q.Add(cur.Left)
			q.Add(cur.Right)

		}
	}

	return true
}

func main() {

}
