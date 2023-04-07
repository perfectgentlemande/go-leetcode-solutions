package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	cur := root

	for cur != nil {
		if cur.Val == val {
			return cur
		}

		if val > cur.Val {
			cur = cur.Right
		} else if val < cur.Val {
			cur = cur.Left
		}
	}

	return nil
}

func main() {

}
