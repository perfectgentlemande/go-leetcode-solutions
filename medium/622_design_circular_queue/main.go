package main

type MyCircularQueue struct {
	vals   []int
	length int
	l      int
	r      int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		vals:   make([]int, k),
		length: 0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	} else if this.IsEmpty() {
		this.vals[this.l] = value
		this.length++
	} else {
		this.r = (this.r + 1) % len(this.vals)
		this.vals[this.r] = value
		this.length++
	}

	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.vals[this.l] = -1
	if this.l != this.r {
		this.l = (this.l + 1) % len(this.vals)
	}
	this.length--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.vals[this.l]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.vals[this.r]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.length == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.length == len(this.vals)
}

func main() {

}
