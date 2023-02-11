package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := maxDepth(root.Left)
	leftRight := maxDepth(root.Right)

	return max(leftHeight, leftRight) + 1
}

type Deque []TreeNodeLevel

func (d Deque) IsEmpty() bool {
	return len(d) == 0
}
func (d *Deque) PushBack(node TreeNodeLevel) {
	*d = append(*d, node)
}
func (d *Deque) PushFront(node TreeNodeLevel) {
	*d = append([]TreeNodeLevel{node}, *d...)
}
func (d *Deque) PopBack() (TreeNodeLevel, bool) {
	if d.IsEmpty() {
		return TreeNodeLevel{}, false
	} else {
		index := len(*d) - 1
		element := (*d)[index]
		*d = (*d)[:index]
		return element, true
	}
}
func (d *Deque) PopFront() (TreeNodeLevel, bool) {
	if d.IsEmpty() {
		return TreeNodeLevel{}, false
	} else {
		element := (*d)[0]
		*d = (*d)[1:]
		return element, true
	}
}

type TreeNodeLevel struct {
	TreeNode
	Level int
}

func maxDepthIterative(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 0
	deque := Deque{
		{
			TreeNode: *root,
			Level:    1,
		},
	}

	for !deque.IsEmpty() {
		curNodeLevel, _ := deque.PopFront()
		if curNodeLevel.Level > res {
			res = curNodeLevel.Level
		}

		if curNodeLevel.Left != nil {
			deque.PushBack(TreeNodeLevel{
				TreeNode: *curNodeLevel.Left,
				Level:    curNodeLevel.Level + 1,
			})
		}
		if curNodeLevel.Right != nil {
			deque.PushBack(TreeNodeLevel{
				TreeNode: *curNodeLevel.Right,
				Level:    curNodeLevel.Level + 1,
			})
		}
	}

	return res
}

func main() {

}
