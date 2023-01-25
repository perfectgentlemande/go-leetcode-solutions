package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	left := postorderTraversal(root.Left)
	right := postorderTraversal(root.Right)

	output := make([]int, 0)

	output = append(output, left...)
	output = append(output, right...)
	output = append(output, root.Val)

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

func reverseSlice(s []int) []int {
	l, r := 0, len(s)-1

	for l <= r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}

	return s
}
func postorderTraversalIterative(root *TreeNode) []int {
	res := make([]int, 0)
	s := make(Stack, 0)

	if root == nil {
		return nil
	}
	s.Push(*root)

	for !s.IsEmpty() {
		node, _ := s.Pop()
		res = append(res, node.Val)

		if node.Left != nil {
			s.Push(*node.Left)
		}
		if node.Right != nil {
			s.Push(*node.Right)
		}
	}

	return reverseSlice(res)
}

func main() {
	root := TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	output := postorderTraversal(&root)
	fmt.Println(output)

	output = postorderTraversal(nil)
	fmt.Println(output)

	root = TreeNode{Val: 1}
	output = postorderTraversal(&root)
	fmt.Println(output)
}
