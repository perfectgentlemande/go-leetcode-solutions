package main

type StackNode struct {
	Val  string
	Next *StackNode
}
type Stack struct {
	root *StackNode
	size int
}

func (s *Stack) Push(val string) {
	nextRoot := s.root

	s.root = &StackNode{
		Val:  val,
		Next: nextRoot,
	}
	s.size++
}

func (s *Stack) Pop() (string, bool) {
	if s.root == nil {
		return "", false
	}

	top := s.root.Val
	s.root = s.root.Next
	s.size--

	return top, true
}

func (s *Stack) Top() (string, bool) {
	if s.root == nil {
		return "", false
	}

	return s.root.Val, true
}

func (s *Stack) IsEmpty() bool {
	return s.root == nil
}

func (s *Stack) Size() int {
	return s.size
}

type BrowserHistory struct {
	BackStack    Stack
	ForwardStack Stack
}

func Constructor(homepage string) BrowserHistory {
	bh := BrowserHistory{}
	bh.BackStack.Push(homepage)

	return bh
}

func (this *BrowserHistory) Visit(url string) {
	this.BackStack.Push(url)
	this.ForwardStack.root = nil
}

func (this *BrowserHistory) Back(steps int) string {
	for i := 0; i < steps && this.BackStack.Size() > 1; i++ {
		v, _ := this.BackStack.Pop()
		this.ForwardStack.Push(v)
	}

	v, _ := this.BackStack.Top()
	return v
}

func (this *BrowserHistory) Forward(steps int) string {
	for i := 0; i < steps && !this.ForwardStack.IsEmpty(); i++ {
		v, _ := this.ForwardStack.Pop()
		this.BackStack.Push(v)
	}

	v, _ := this.BackStack.Top()
	return v
}

func main() {

}
