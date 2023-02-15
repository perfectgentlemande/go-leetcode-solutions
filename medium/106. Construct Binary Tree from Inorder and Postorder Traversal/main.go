package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NOTE: I don't understand how it works: I just re-wrote the official solution into Go and it works :(

func helper(inLeft, inRight int, postorder []int, postIdx *int, idxMap map[int]int) *TreeNode {
	if inLeft > inRight {
		return nil
	}

	rootVal := postorder[*postIdx]
	root := TreeNode{Val: rootVal}
	index := idxMap[rootVal]
	*postIdx = (*postIdx) - 1

	root.Right = helper(index+1, inRight, postorder, postIdx, idxMap)
	root.Left = helper(inLeft, index-1, postorder, postIdx, idxMap)

	return &root
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	postIdx := len(postorder) - 1
	idx := 0
	idxMap := map[int]int{}

	for _, val := range inorder {
		idxMap[val] = idx
		idx++
	}

	return helper(0, len(inorder)-1, postorder, &postIdx, idxMap)
}

func main() {

}
