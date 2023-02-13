package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum -= root.Val
	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}

	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}

type Deque []TreeNodeSum

func (d Deque) IsEmpty() bool {
	return len(d) == 0
}
func (d *Deque) PushBack(node TreeNodeSum) {
	*d = append(*d, node)
}
func (d *Deque) PushFront(node TreeNodeSum) {
	*d = append([]TreeNodeSum{node}, *d...)
}
func (d *Deque) PopBack() (TreeNodeSum, bool) {
	if d.IsEmpty() {
		return TreeNodeSum{}, false
	} else {
		index := len(*d) - 1
		element := (*d)[index]
		*d = (*d)[:index]
		return element, true
	}
}
func (d *Deque) PopFront() (TreeNodeSum, bool) {
	if d.IsEmpty() {
		return TreeNodeSum{}, false
	} else {
		element := (*d)[0]
		*d = (*d)[1:]
		return element, true
	}
}

type TreeNodeSum struct {
	TreeNode
	Sum int
}

func hasPathSumIterative(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	deque := Deque{TreeNodeSum{
		TreeNode: *root,
		Sum:      targetSum,
	}}

	for !deque.IsEmpty() {
		curNodeSum, _ := deque.PopFront()

		if curNodeSum.Sum-curNodeSum.Val == 0 &&
			curNodeSum.Left == nil && curNodeSum.Right == nil {
			return true
		}

		if curNodeSum.Left != nil {
			deque.PushBack(TreeNodeSum{
				TreeNode: *curNodeSum.Left,
				Sum:      curNodeSum.Sum - curNodeSum.Val,
			})
		}
		if curNodeSum.Right != nil {
			deque.PushBack(TreeNodeSum{
				TreeNode: *curNodeSum.Right,
				Sum:      curNodeSum.Sum - curNodeSum.Val,
			})
		}
	}

	return false
}

func main() {

}
