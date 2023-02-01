package main

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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return false
	}

	if p != nil && q == nil {
		return false
	}

	if q != nil && p == nil {
		return false
	}

	return (p.Val == q.Val) && isSameTree(p.Left, q.Left) && isSameTree(p.Right, p.Right)
}

func main() {

}
