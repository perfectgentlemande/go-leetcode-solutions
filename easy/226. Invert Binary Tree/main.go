package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	invertedLeft := invertTree(root.Left)
	invertedRight := invertTree(root.Right)

	root.Left = invertedRight
	root.Right = invertedLeft

	return root
}

type Queue []*TreeNode

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}
func (q *Queue) Add(node *TreeNode) {
	*q = append([]*TreeNode{node}, *q...)
}
func (q *Queue) Remove() (*TreeNode, bool) {
	if q.IsEmpty() {
		return &TreeNode{}, false
	} else {
		index := len(*q) - 1
		element := (*q)[index]
		*q = (*q)[:index]
		return element, true
	}
}

func invertTreeIterative(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := Queue{root}

	for !queue.IsEmpty() {
		item, _ := queue.Remove()

		item.Left, item.Right = item.Right, item.Left
		if item.Left != nil {
			queue.Add(item.Left)
		}

		if item.Right != nil {
			queue.Add(item.Right)
		}
	}

	return root
}

func main() {

}
