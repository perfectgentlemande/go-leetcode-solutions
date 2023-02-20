package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// got it here, I'm tired of trees and still noob so I found this there:
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/solutions/213243/simple-and-efficent-golang-100/?q=golang&orderBy=most_relevant

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	lc := lowestCommonAncestor(root.Left, p, q)
	rc := lowestCommonAncestor(root.Right, p, q)
	if lc == nil {
		return rc
	}
	if rc == nil {
		return lc
	}
	return root
}

func main() {

}
