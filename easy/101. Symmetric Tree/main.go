package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isMirror(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if (node1 == nil) != (node2 == nil) {
		return false
	}

	return (node1.Val == node2.Val) &&
		isMirror(node1.Left, node2.Right) &&
		isMirror(node1.Right, node2.Left)
}
func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

type Deque []TreeNode

func (d Deque) IsEmpty() bool {
	return len(d) == 0
}
func (d *Deque) PushBack(node TreeNode) {
	*d = append(*d, node)
}
func (d *Deque) PushFront(node TreeNode) {
	*d = append([]TreeNode{node}, *d...)
}
func (d *Deque) PopBack() (TreeNode, bool) {
	if d.IsEmpty() {
		return TreeNode{}, false
	} else {
		index := len(*d) - 1
		element := (*d)[index]
		*d = (*d)[:index]
		return element, true
	}
}
func (d *Deque) PopFront() (TreeNode, bool) {
	if d.IsEmpty() {
		return TreeNode{}, false
	} else {
		element := (*d)[0]
		*d = (*d)[1:]
		return element, true
	}
}

func isSymmetricIterative(root *TreeNode) bool {
	if root == nil {
		return true
	}

	deque1 := Deque{*root}
	deque2 := Deque{*root}

	for !deque1.IsEmpty() {
		d1Node, _ := deque1.PopFront()
		d2Node, _ := deque2.PopFront()

		if d1Node.Val != d2Node.Val {
			return false
		}

		if (d1Node.Left == nil) != (d2Node.Right == nil) {
			return false
		}
		if (d1Node.Right == nil) != (d2Node.Left == nil) {
			return false
		}

		if d1Node.Left != nil {
			deque1.PushBack(*d1Node.Left)
		}
		if d1Node.Right != nil {
			deque1.PushBack(*d1Node.Right)
		}
		if d2Node.Right != nil {
			deque2.PushBack(*d2Node.Right)
		}
		if d2Node.Left != nil {
			deque2.PushBack(*d2Node.Left)
		}
	}

	return true
}

func main() {

}
