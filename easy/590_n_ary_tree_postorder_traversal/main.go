package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}

	output := make([]int, 0)

	for i := range root.Children {
		output = append(output, postorder(root.Children[i])...)
	}

	output = append(output, root.Val)

	return output
}

type Stack []Node

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func (s *Stack) Push(node Node) {
	*s = append(*s, node)
}

func (s *Stack) Pop() (Node, bool) {
	if s.IsEmpty() {
		return Node{}, false
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

func postorderIterative(root *Node) []int {
	res := make([]int, 0)
	s := make(Stack, 0)

	if root == nil {
		return nil
	}
	s.Push(*root)

	for !s.IsEmpty() {
		node, _ := s.Pop()
		res = append(res, node.Val)

		for i := range node.Children {
			if node.Children[i] != nil {
				s.Push(*node.Children[i])
			}
		}
	}

	return reverseSlice(res)
}

func main() {
	root := Node{Val: 1}
	root.Children = []*Node{
		{
			Val: 3,
			Children: []*Node{
				{
					Val: 5,
				},
				{
					Val: 6,
				},
			},
		},
		{
			Val: 2,
		},
		{
			Val: 4,
		},
	}
	output := postorder(&root)
	fmt.Println(output)
}
