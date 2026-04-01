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

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	deque := Deque{TreeNodeLevel{TreeNode: *root, Level: 0}}

	for !deque.IsEmpty() {
		curItem, _ := deque.PopFront()

		if curItem.Level == len(res) {
			res = append(res, make([]int, 0))
		}

		fmt.Println(curItem.Level)

		if (curItem.Level+1)%2 == 0 {
			if curItem.Left != nil {
				deque.PushFront(TreeNodeLevel{TreeNode: *curItem.Left, Level: curItem.Level + 1})
			}

			if curItem.Right != nil {
				deque.PushFront(TreeNodeLevel{TreeNode: *curItem.Right, Level: curItem.Level + 1})
			}

			res[curItem.Level] = append(res[curItem.Level], curItem.Val)
		} else {
			if curItem.Right != nil {
				deque.PushBack(TreeNodeLevel{TreeNode: *curItem.Right, Level: curItem.Level + 1})
			}

			if curItem.Left != nil {
				deque.PushBack(TreeNodeLevel{TreeNode: *curItem.Left, Level: curItem.Level + 1})
			}

			res[curItem.Level] = append([]int{curItem.Val}, res[curItem.Level]...)
		}
	}

	return res
}

func main() {

}
