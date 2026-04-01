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

func searchBSTRecursive(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}

	if val > root.Val {
		return searchBSTRecursive(root.Right, val)
	} else if val < root.Val {
		return searchBSTRecursive(root.Left, val)
	}

	return nil
}

func main() {

}
