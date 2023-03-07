package main

type Node struct {
	Val  int
	Next *Node
}
type MinStack struct {
	root    *Node
	minRoot *Node
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	nextRoot := this.root
	nextMinRoot := &Node{
		Next: this.minRoot,
	}

	if this.minRoot == nil || val < this.minRoot.Val {
		nextMinRoot.Val = val
	} else {
		nextMinRoot.Val = this.minRoot.Val
	}

	this.minRoot = nextMinRoot
	this.root = &Node{
		Val:  val,
		Next: nextRoot,
	}
}

func (this *MinStack) Pop() {
	if this.root == nil {
		return
	}

	this.root = this.root.Next
	this.minRoot = this.minRoot.Next
}

func (this *MinStack) Top() int {
	if this.root == nil {
		return 0
	}

	return this.root.Val
}

func (this *MinStack) GetMin() int {
	if this.root == nil {
		return 0
	}

	return this.minRoot.Val
}

func main() {

}
