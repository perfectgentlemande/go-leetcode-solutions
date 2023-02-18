package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type Queue []*Node

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}
func (q *Queue) Add(node *Node) {
	*q = append([]*Node{node}, *q...)
}
func (q *Queue) Remove() (*Node, bool) {
	if q.IsEmpty() {
		return &Node{}, false
	} else {
		index := len(*q) - 1
		element := (*q)[index]
		*q = (*q)[:index]
		return element, true
	}
}

func levelOrder(root *Node) [][]*Node {
	if root == nil {
		return nil
	}

	res := make([][]*Node, 0)
	level := 0
	queue := Queue{}
	queue.Add(root)

	for !queue.IsEmpty() {
		res = append(res, make([]*Node, 0))
		levelLen := len(queue)

		for i := 0; i < levelLen; i++ {
			node, _ := queue.Remove()
			res[level] = append(res[level], node)

			if node.Left != nil {
				queue.Add(node.Left)
			}
			if node.Right != nil {
				queue.Add(node.Right)
			}
		}

		level++
	}

	return res
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	lot := levelOrder(root)
	for i := range lot {
		for j := 0; j < len(lot[i])-1; j++ {
			curNode := (*lot[i][j])
			curNode.Next = lot[i][j+1]
			*lot[i][j] = curNode
		}
	}

	return root
}

func main() {

}
