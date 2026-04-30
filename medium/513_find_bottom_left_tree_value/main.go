package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue []TreeNode

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}
func (q *Queue) Add(node TreeNode) {
	*q = append([]TreeNode{node}, *q...)
}
func (q *Queue) Remove() (TreeNode, bool) {
	if q.IsEmpty() {
		return TreeNode{}, false
	} else {
		index := len(*q) - 1
		element := (*q)[index]
		*q = (*q)[:index]
		return element, true
	}
}

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return -1
	}

	res := root.Val
	queue := Queue{}
	queue.Add(*root)

	for !queue.IsEmpty() {
		cur, _ := queue.Remove()
		res = cur.Val

		if cur.Right != nil {
			queue.Add(*cur.Right)
		}
		if cur.Left != nil {
			queue.Add(*cur.Left)
		}
	}

	return res
}

func main() {

}
