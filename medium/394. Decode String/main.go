package main

import (
	"unicode"
)

type StackItem interface {
	int | string
}

type StackNode[T StackItem] struct {
	Val  T
	Next *StackNode[T]
}
type Stack[T StackItem] struct {
	root *StackNode[T]
}

func (s *Stack[T]) Push(val T) {
	nextRoot := s.root

	s.root = &StackNode[T]{
		Val:  val,
		Next: nextRoot,
	}
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.root == nil {
		n := StackNode[T]{}

		return n.Val, false
	}

	curVal := s.root.Val
	s.root = s.root.Next

	return curVal, true
}

func (s *Stack[T]) Top() (T, bool) {
	if s.root == nil {
		n := StackNode[T]{}

		return n.Val, false
	}

	return s.root.Val, true
}

func (s *Stack[T]) IsEmpty() bool {
	return s.root == nil
}
func decodeString(s string) string {
	curString := ""
	curNum := 0

	numStack := Stack[int]{}
	strStack := Stack[string]{}

	for _, rn := range s {
		if rn == '[' {
			strStack.Push(curString)
			numStack.Push(curNum)
			curString = ""
			curNum = 0
		} else if rn == ']' {
			prevStrV, _ := strStack.Pop()
			prevNumV, _ := numStack.Pop()

			newStr := ""
			for i := 0; i < prevNumV; i++ {
				newStr += curString
			}
			curString = prevStrV + newStr
		} else if unicode.IsDigit(rn) {
			curNum = curNum*10 + int(rn) - '0'
		} else {
			curString += string(rn)
		}
	}

	return curString
}

type StackStringNode struct {
	Val  string
	Next *StackStringNode
}
type StackString struct {
	root *StackStringNode
}

func (s *StackString) Push(val string) {
	nextRoot := s.root

	s.root = &StackStringNode{
		Val:  val,
		Next: nextRoot,
	}
}

func (s *StackString) Pop() (string, bool) {
	if s.root == nil {
		return "", false
	}

	curVal := s.root.Val
	s.root = s.root.Next

	return curVal, true
}

func (s *StackString) Top() (string, bool) {
	if s.root == nil {
		return "", false
	}

	return s.root.Val, true
}

func (s *StackString) IsEmpty() bool {
	return s.root == nil
}

type StackIntNode struct {
	Val  int
	Next *StackIntNode
}
type StackInt struct {
	root *StackIntNode
}

func (s *StackInt) Push(val int) {
	nextRoot := s.root

	s.root = &StackIntNode{
		Val:  val,
		Next: nextRoot,
	}
}

func (s *StackInt) Pop() (int, bool) {
	if s.root == nil {
		return 0, false
	}

	curVal := s.root.Val
	s.root = s.root.Next

	return curVal, true
}

func (s *StackInt) Top() (int, bool) {
	if s.root == nil {
		return 0, false
	}

	return s.root.Val, true
}

func (s *StackInt) IsEmpty() bool {
	return s.root == nil
}

func decodeStringNoGenerics(s string) string {
	curString := ""
	curNum := 0

	numStack := StackInt{}
	strStack := StackString{}

	for _, rn := range s {
		if rn == '[' {
			strStack.Push(curString)
			numStack.Push(curNum)
			curString = ""
			curNum = 0
		} else if rn == ']' {
			prevStrV, _ := strStack.Pop()
			prevNumV, _ := numStack.Pop()

			newStr := ""
			for i := 0; i < prevNumV; i++ {
				newStr += curString
			}
			curString = prevStrV + newStr
		} else if unicode.IsDigit(rn) {
			curNum = curNum*10 + int(rn) - '0'
		} else {
			curString += string(rn)
		}
	}

	return curString
}
func main() {

}
