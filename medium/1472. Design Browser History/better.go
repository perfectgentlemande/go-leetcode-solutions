package main

type DLLNode struct {
	Val      string
	Next     *DLLNode
	Previous *DLLNode
}

type BrowserHistory2 struct {
	head    *DLLNode
	current *DLLNode
}

func Constructor(homepage string) BrowserHistory2 {
	bh := BrowserHistory2{
		head: &DLLNode{
			Val: homepage,
		},
	}

	bh.current = bh.head

	return bh
}

func (this *BrowserHistory2) Visit(url string) {
	newNode := &DLLNode{
		Val: url,
	}

	this.current.Next = newNode
	newNode.Previous = this.current
	this.current = newNode
}

func (this *BrowserHistory2) Back(steps int) string {
	for i := steps; i > 0 && (this.current.Previous != nil); i-- {
		this.current = this.current.Previous
	}

	return this.current.Val
}

func (this *BrowserHistory2) Forward(steps int) string {
	for i := steps; i > 0 && (this.current.Next != nil); i-- {
		this.current = this.current.Next
	}

	return this.current.Val
}
