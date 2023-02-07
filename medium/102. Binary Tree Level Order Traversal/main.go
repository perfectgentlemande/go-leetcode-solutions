package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeLevel struct {
	TreeNode
	Level int
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

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	deque := Deque{}
	deque.PushBack(TreeNodeLevel{
		TreeNode: *root,
		Level:    0,
	})
	res = append(res, make([]int, 0))

	for !deque.IsEmpty() {
		curNode, _ := deque.PopBack()
		res[curNode.Level] = append(res[curNode.Level], curNode.Val)

		if curNode.Left == nil && curNode.Right == nil {
			continue
		}

		if len(res) == curNode.Level+1 {
			res = append(res, make([]int, 0))
		}

		if curNode.Left != nil {
			deque.PushFront(TreeNodeLevel{
				TreeNode: *curNode.Left,
				Level:    curNode.Level + 1,
			})
		}
		if curNode.Right != nil {
			deque.PushFront(TreeNodeLevel{
				TreeNode: *curNode.Right,
				Level:    curNode.Level + 1,
			})
		}
	}

	return res
}

func main() {
	r := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(levelOrder(&r))

	r = TreeNode{
		Val: 1,
	}
	fmt.Println(levelOrder(&r))
	fmt.Println(levelOrder(nil))

	r = TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	fmt.Println(levelOrder(&r))
}
