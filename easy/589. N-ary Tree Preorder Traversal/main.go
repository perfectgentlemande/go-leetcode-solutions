package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	output := make([]int, 0)

	output = append(output, root.Val)

	for i := range root.Children {
		output = append(output, preorder(root.Children[i])...)
	}

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

func preorderIterative(root *Node) []int {
	res := make([]int, 0)
	s := make(Stack, 0)

	if root != nil {
		s.Push(*root)
	}

	var cur Node

	for !s.IsEmpty() {
		cur, _ = s.Pop()

		res = append(res, cur.Val)
		for i := len(cur.Children) - 1; i >= 0; i-- {
			if cur.Children[i] != nil {
				s.Push(*cur.Children[i])
			}
		}
	}

	return res
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
	output := preorder(&root)
	fmt.Println(output)

	root = Node{Val: 1}
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
	output = preorderIterative(&root)
	fmt.Println(output)
}
