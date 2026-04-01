package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)

	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minDiffInBST(root *TreeNode) int {
	minDistance := math.MaxInt64

	vals := inorderTraversal(root)
	for i := 1; i < len(vals); i++ {
		minDistance = min(minDistance, vals[i]-vals[i-1])
	}

	return minDistance
}

func main() {

}
