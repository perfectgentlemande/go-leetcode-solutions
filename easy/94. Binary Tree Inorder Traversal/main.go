package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)

	output := make([]int, 0)

	output = append(output, left...)
	output = append(output, root.Val)
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

func inorderTraversalIterative(root *TreeNode) []int {
	res := make([]int, 0)
	s := make(Stack, 0)

	cur := root
	for cur != nil || !s.IsEmpty() {
		for cur != nil {
			s.Push(*cur)
			cur = cur.Left
		}

		newCur, _ := s.Pop()
		cur = &newCur
		res = append(res, cur.Val)
		cur = cur.Right
	}

	return res
}

func main() {
	root := TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	output := inorderTraversal(&root)
	fmt.Println(output)

	output = inorderTraversal(nil)
	fmt.Println(output)

	root = TreeNode{Val: 1}
	output = inorderTraversal(&root)
	fmt.Println(output)

	root = TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	output = inorderTraversalIterative(&root)
	fmt.Println(output)

	output = inorderTraversalIterative(nil)
	fmt.Println(output)

	root = TreeNode{Val: 1}
	output = inorderTraversalIterative(&root)
	fmt.Println(output)
}
