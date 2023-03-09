package main

type Node struct {
	Val       int
	Neighbors []*Node
}

type StackNode struct {
	Val  Node
	Next *StackNode
}
type Stack struct {
	root *StackNode
}

func (s *Stack) Push(val Node) {
	nextRoot := s.root

	s.root = &StackNode{
		Val:  val,
		Next: nextRoot,
	}
}

func (s *Stack) Pop() {
	if s.root == nil {
		return
	}

	s.root = s.root.Next
}

func (s *Stack) Top() Node {
	if s.root == nil {
		return Node{}
	}

	return s.root.Val
}

func (s *Stack) IsEmpty() bool {
	return s.root == nil
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}

	s := Stack{}
	s.Push(*node)
	visited := map[int]*Node{}

	for !s.IsEmpty() {
		cur := s.Top()
		s.Pop()

		if _, ok := visited[cur.Val]; !ok {
			visited[cur.Val] = &Node{
				Val: cur.Val,
			}
		}

		for i := range cur.Neighbors {
			if _, ok := visited[cur.Neighbors[i].Val]; !ok {
				visited[cur.Neighbors[i].Val] = &Node{
					Val: cur.Neighbors[i].Val,
				}
				s.Push(*cur.Neighbors[i])
			}

			newCurVisited := visited[cur.Val]
			newCurVisited.Neighbors = append(newCurVisited.Neighbors, visited[cur.Neighbors[i].Val])
		}
	}

	return visited[node.Val]
}

func main() {

}
