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

func traverse(node TreeNode, level int, res *[][]int) {
	if len(*res) == level {
		*res = append(*res, make([]int, 0))
	}
	(*res)[level] = append((*res)[level], node.Val)

	if node.Left != nil {
		traverse(*node.Left, level+1, res)
	}
	if node.Right != nil {
		traverse(*node.Right, level+1, res)
	}
}

func levelOrderRecursive(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	traverse(*root, 0, &res)

	return res
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

func levelOrderIterativeOneMore(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	level := 0
	queue := Queue{}
	queue.Add(*root)

	for !queue.IsEmpty() {
		res = append(res, make([]int, 0))
		levelLen := len(queue)

		for i := 0; i < levelLen; i++ {
			node, _ := queue.Remove()
			res[level] = append(res[level], node.Val)

			if node.Left != nil {
				queue.Add(*node.Left)
			}
			if node.Right != nil {
				queue.Add(*node.Right)
			}
		}

		level++
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

	fmt.Println("recursive")

	r = TreeNode{
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
	fmt.Println(levelOrderRecursive(&r))

	r = TreeNode{
		Val: 1,
	}
	fmt.Println(levelOrderRecursive(&r))
	fmt.Println(levelOrderRecursive(nil))

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
	fmt.Println(levelOrderRecursive(&r))

	fmt.Println("one more iterative")

	r = TreeNode{
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
	fmt.Println(levelOrderIterativeOneMore(&r))

	r = TreeNode{
		Val: 1,
	}
	fmt.Println(levelOrderIterativeOneMore(&r))
	fmt.Println(levelOrderIterativeOneMore(nil))

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
	fmt.Println(levelOrderIterativeOneMore(&r))
}
