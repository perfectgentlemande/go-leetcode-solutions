package main

import "fmt"

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

type MyQueue struct {
	writer Stack
	reader Stack
}

func (this *MyQueue) moveFromWriterToreader() {
	for !this.writer.IsEmpty() {
		v, _ := this.writer.Pop()
		this.reader.Push(v)
	}
}

func (this *MyQueue) returnToWriter() {
	for !this.reader.IsEmpty() {
		v, _ := this.reader.Pop()
		this.writer.Push(v)
	}
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.writer.Push(x)
}

func (this *MyQueue) Pop() int {
	if this.reader.IsEmpty() {
		this.moveFromWriterToreader()
	}

	v, _ := this.reader.Pop()
	return v
}

func (this *MyQueue) Peek() int {
	if this.reader.IsEmpty() {
		this.moveFromWriterToreader()
	}

	v, _ := this.reader.Top()
	return v
}

func (this *MyQueue) Empty() bool {
	return this.reader.IsEmpty() && this.writer.IsEmpty()
}

func main() {
	myQueue := MyQueue{}
	myQueue.Push(1)
	fmt.Println(myQueue.Peek())
	myQueue.Push(2)
	fmt.Println(myQueue.Peek())
	myQueue.Peek()
	myQueue.Pop()
	fmt.Println(myQueue.Peek())
	fmt.Println(myQueue.Empty())
}
