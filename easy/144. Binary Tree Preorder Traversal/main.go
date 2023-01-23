package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	left := preorderTraversal(root.Left)
	right := preorderTraversal(root.Right)

	output := make([]int, 0)

	output = append(output, root.Val)
	output = append(output, left...)
	output = append(output, right...)
	return output
}

type Stack []TreeNode

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func (s *Stack) Push(node TreeNode) {
	*s = append(*s, node)
}

func (s *Stack) Pop() (TreeNode, bool) {
	if s.IsEmpty() {
		return TreeNode{}, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func preorderTraversalIterative(root *TreeNode) []int {
	res := make([]int, 0)
	s := make(Stack, 0)

	if root != nil {
		s.Push(*root)
	}

	var cur TreeNode

	for !s.IsEmpty() {
		cur, _ = s.Pop()

		res = append(res, cur.Val)
		if cur.Right != nil {
			s.Push(*cur.Right)
		}
		if cur.Left != nil {
			s.Push(*cur.Left)
		}
	}

	return res
}

func main() {
	root := TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	output := preorderTraversal(&root)
	fmt.Println(output)

	root = TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	output = preorderTraversal(&root)
	fmt.Println(output)

	output = preorderTraversal(nil)
	fmt.Println(output)

	root = TreeNode{Val: 1}
	output = preorderTraversal(&root)
	fmt.Println(output)

	root = TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	output = preorderTraversalIterative(&root)
	fmt.Println(output)

	root = TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	output = preorderTraversalIterative(&root)
	fmt.Println(output)

	output = preorderTraversalIterative(nil)
	fmt.Println(output)

	root = TreeNode{Val: 1}
	output = preorderTraversalIterative(&root)
	fmt.Println(output)
}
