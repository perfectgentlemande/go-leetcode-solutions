package main

type QueueNode struct {
	Val  int
	Next *QueueNode
}
type Queue struct {
	root *QueueNode
	rear *QueueNode
	size int
}

func (q *Queue) Add(val int) {
	nodePtr := &QueueNode{
		Val: val,
	}

	if q.root == nil {
		q.root = nodePtr
	}

	if q.rear != nil {
		q.rear.Next = nodePtr
	}

	q.rear = nodePtr
	q.size++
}

func (q *Queue) Remove() (int, bool) {
	if q.root == nil {
		return 0, false
	}

	curVal := q.root.Val
	if q.root == q.rear {
		q.rear = nil
	}
	q.root = q.root.Next
	q.size--

	return curVal, true
}

func (q *Queue) Head() (int, bool) {
	if q.root == nil {
		return 0, false
	}

	return q.root.Val, true
}

func (q *Queue) IsEmpty() bool {
	return q.root == nil
}

func (q *Queue) Size() int {
	return q.size
}

type MyStack struct {
	q1 Queue
	q2 Queue
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.q1.Add(x)

	for this.q1.Size() > 1 {
		v, _ := this.q1.Remove()

		this.q2.Add(v)
	}
}

func (this *MyStack) Pop() int {
	if this.q1.Size() == 0 {
		return 0
	}

	popped, _ := this.q1.Remove()

	for this.q2.Size() > 1 {
		v, _ := this.q2.Remove()

		this.q1.Add(v)
	}

	this.q1, this.q2 = this.q2, this.q1
	return popped
}

func (this *MyStack) Top() int {
	if this.q1.size == 0 {
		return 0
	}

	v, _ := this.q1.Head()
	return v
}

func (this *MyStack) Empty() bool {
	return this.q1.IsEmpty() && this.q2.IsEmpty()
}

func main() {

}
