package main

type TempIndex struct {
	Temp  int
	Index int
}
type Node struct {
	Val  TempIndex
	Next *Node
}
type Stack struct {
	root *Node
}

func Constructor() Stack {
	return Stack{}
}

func (s *Stack) Push(val TempIndex) {
	nextRoot := s.root

	s.root = &Node{
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

func (s *Stack) Top() TempIndex {
	if s.root == nil {
		return TempIndex{}
	}

	return s.root.Val
}

func (s *Stack) IsEmpty() bool {
	return s.root == nil
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	s := Stack{}

	for i, t := range temperatures {
		for !s.IsEmpty() && s.Top().Temp < t {
			top := s.Top()
			res[top.Index] = i - top.Index
			s.Pop()
		}

		s.Push(TempIndex{Temp: t, Index: i})
	}

	return res
}

func main() {

}
