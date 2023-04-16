package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	nodes := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		nodes = append(nodes, i)
	}

	return buildTree(nodes)
}

func buildTree(nodes []int) []*TreeNode {
	if len(nodes) == 0 {
		return []*TreeNode{nil}
	}

	res := make([]*TreeNode, 0)

	for i := 0; i < len(nodes); i++ {
		lt := buildTree(nodes[:i])
		rt := buildTree(nodes[i+1:])

		for _, l := range lt {
			for _, r := range rt {
				res = append(res, &TreeNode{
					Val:   nodes[i],
					Left:  l,
					Right: r,
				})
			}
		}
	}

	return res
}
