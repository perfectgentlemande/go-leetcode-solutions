package main

import "strconv"

type Node struct {
	Val  int
	Next *Node
}
type Stack struct {
	root *Node
}

func Constructor() Stack {
	return Stack{}
}

func (s *Stack) Push(val int) {
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

func (s *Stack) Top() int {
	if s.root == nil {
		return 0
	}

	return s.root.Val
}

func (s *Stack) IsEmpty() bool {
	return s.root == nil
}

func getOperation(operation string) (func(a, b int) int, bool) {
	v, ok := map[string]func(a, b int) int{
		"+": func(a, b int) int {
			return a + b
		},
		"-": func(a, b int) int {
			return a - b
		},
		"/": func(a, b int) int {
			return a / b
		},
		"*": func(a, b int) int {
			return a * b
		},
	}[operation]

	return v, ok
}

func evalRPN(tokens []string) int {
	s := Stack{}

	for i := 0; i < len(tokens); i++ {
		operation, ok := getOperation(tokens[i])
		if !ok {
			val, _ := strconv.Atoi(tokens[i])
			s.Push(val)
			continue
		}

		b := s.Top()
		s.Pop()
		a := s.Top()
		s.Pop()

		s.Push(operation(a, b))
	}

	return s.Top()
}
func main() {

}
