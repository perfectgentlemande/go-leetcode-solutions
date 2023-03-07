package main

type Node struct {
	Val  int
	Next *Node
}
type MinStack struct {
	root *Node
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	nextRoot := this.root

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

	min := this.root.Val
	cur := this.root

	for cur != nil {
		if cur.Val < min {
			min = cur.Val
		}

		cur = cur.Next
	}

	return min
}

func main() {

}
