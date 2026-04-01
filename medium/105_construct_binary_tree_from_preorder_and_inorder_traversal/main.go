package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NOTE: I don't understand how it works: I just re-wrote the official solution into Go and it works :(

func helper(inLeft, inRight int, preorder []int, preIdx *int, idxMap map[int]int) *TreeNode {
	if inLeft > inRight {
		return nil
	}

	rootVal := preorder[*preIdx]
	root := TreeNode{Val: rootVal}
	*preIdx = (*preIdx) + 1

	root.Left = helper(inLeft, idxMap[rootVal]-1, preorder, preIdx, idxMap)
	root.Right = helper(idxMap[rootVal]+1, inRight, preorder, preIdx, idxMap)

	return &root
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	preIdx := 0
	idxMap := map[int]int{}

	for i, val := range inorder {
		idxMap[val] = i
	}

	return helper(0, len(inorder)-1, preorder, &preIdx, idxMap)
}

func main() {
}
