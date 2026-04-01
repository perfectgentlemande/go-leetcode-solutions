package main

type Node struct {
	Val      int
	Children []*Node
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	res := 0
	for i := range root.Children {
		res = max(res, maxDepth(root.Children[i]))
	}

	return res + 1
}

type NodeLevel struct {
	Node
	Level int
}

type Deque []NodeLevel

func (d Deque) IsEmpty() bool {
	return len(d) == 0
}
func (d *Deque) PushBack(node NodeLevel) {
	*d = append(*d, node)
}
func (d *Deque) PushFront(node NodeLevel) {
	*d = append([]NodeLevel{node}, *d...)
}
func (d *Deque) PopBack() (NodeLevel, bool) {
	if d.IsEmpty() {
		return NodeLevel{}, false
	} else {
		index := len(*d) - 1
		element := (*d)[index]
		*d = (*d)[:index]
		return element, true
	}
}
func (d *Deque) PopFront() (NodeLevel, bool) {
	if d.IsEmpty() {
		return NodeLevel{}, false
	} else {
		element := (*d)[0]
		*d = (*d)[1:]
		return element, true
	}
}

func maxDepthIterative(root *Node) int {
	if root == nil {
		return 0
	}

	res := 0
	deque := Deque{
		{
			Node:  *root,
			Level: 1,
		},
	}

	for !deque.IsEmpty() {
		curNodeLevel, _ := deque.PopFront()
		if curNodeLevel.Level > res {
			res = curNodeLevel.Level
		}

		for i := range curNodeLevel.Children {
			if curNodeLevel.Children[i] != nil {
				deque.PushBack(NodeLevel{
					Node:  *curNodeLevel.Children[i],
					Level: curNodeLevel.Level + 1,
				})
			}
		}
	}

	return res
}

func main() {

}
