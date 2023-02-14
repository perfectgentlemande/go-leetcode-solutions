package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUni(node TreeNode, count *int) bool {
	if node.Left == nil && node.Right == nil {
		*count += 1
		return true
	}

	isUnival := true

	if node.Left != nil {
		isUnival = isUni(*node.Left, count) && isUnival && node.Left.Val == node.Val
	}
	if node.Right != nil {
		isUnival = isUni(*node.Right, count) && isUnival && node.Right.Val == node.Val
	}

	if !isUnival {
		return false
	}

	*count += 1
	return true
}

func countUnivalSubtrees(root *TreeNode) int {
	if root == nil {
		return 0
	}

	count := 0
	isUni(*root, &count)

	return count
}

func main() {

}
