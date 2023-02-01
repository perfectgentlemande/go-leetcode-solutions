package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
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
	p := TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	q := TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(isSameTree(&p, &q))

	p = TreeNode{Val: 1,
		Left: &TreeNode{Val: 2},
	}
	q = TreeNode{Val: 1,
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(isSameTree(&p, &q))
}
