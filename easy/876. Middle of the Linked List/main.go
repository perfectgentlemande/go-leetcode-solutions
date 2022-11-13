package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// for debug
func insert(items []int) *ListNode {
	if len(items) == 0 {
		return nil
	}

	firstNode := &ListNode{
		Val:  items[0],
		Next: nil,
	}

	curNode := firstNode
	for i := 1; i < len(items); i++ {
		nextNode := &ListNode{
			Val:  items[i],
			Next: nil,
		}
		curNode.Next = nextNode
		curNode = nextNode
	}

	return firstNode
}
func show(l *ListNode) {
	if l == nil {
		return
	}

	p := l
	for p != nil {
		fmt.Printf("-> %v ", p.Val)
		p = p.Next
	}
}

// solution functions
func count(head *ListNode) int {
	p := head
	count := 0
	for p != nil {
		count++
		p = p.Next
	}

	return count
}
func middleNode(head *ListNode) *ListNode {
	count := count(head)
	if count == 0 {
		return nil
	}
	middleIndex := count / 2

	found := head
	curInd := 0
	for found != nil && curInd < middleIndex {
		curInd++
		found = found.Next
	}

	return found
}

func main() {
	ln1 := insert([]int{1, 2, 3, 4, 5})
	show(ln1)
	fmt.Println(middleNode(ln1).Val)
	ln2 := insert([]int{1, 2, 3, 4, 5, 6})
	show(ln2)
	fmt.Println(middleNode(ln2).Val)
}
