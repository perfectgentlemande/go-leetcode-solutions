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
