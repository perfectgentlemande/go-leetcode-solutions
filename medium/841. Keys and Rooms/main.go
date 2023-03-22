package main

type StackNode struct {
	Val  int
	Next *StackNode
}
type Stack struct {
	root *StackNode
}

func (s *Stack) Push(val int) {
	nextRoot := s.root

	s.root = &StackNode{
		Val:  val,
		Next: nextRoot,
	}
}

func (s *Stack) Pop() (int, bool) {
	if s.root == nil {
		return 0, false
	}

	top := s.root.Val
	s.root = s.root.Next

	return top, true
}

func (s *Stack) Top() (int, bool) {
	if s.root == nil {
		return 0, false
	}

	return s.root.Val, true
}

func (s *Stack) IsEmpty() bool {
	return s.root == nil
}

func canVisitAllRooms(rooms [][]int) bool {
	if len(rooms) == 0 {
		return true
	}

	visited := map[int]struct{}{}

	s := Stack{}
	s.Push(0)

	for !s.IsEmpty() {
		roomNum, _ := s.Pop()

		if _, ok := visited[roomNum]; ok {
			continue
		}

		for _, key := range rooms[roomNum] {
			s.Push(key)
		}

		visited[roomNum] = struct{}{}
	}

	return len(visited) == len(rooms)
}

func main() {

}
