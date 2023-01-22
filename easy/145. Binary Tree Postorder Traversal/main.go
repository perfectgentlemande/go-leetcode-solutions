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
