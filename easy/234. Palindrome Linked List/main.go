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
	fmt.Print("\n")
}

// solution functions
func invert(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	curr := head.Next
	prev := &ListNode{
		Val: head.Val,
	}
	for curr != nil {
		prev = &ListNode{
			Val:  curr.Val,
			Next: prev,
		}
		curr = curr.Next
	}

	return prev
}
func isPalindrome(head *ListNode) bool {
	ass := invert(head)

	p1, p2 := head, ass
	for p1 != nil {
		if p1.Val != p2.Val {
			return false
		}

		p1 = p1.Next
		p2 = p2.Next
	}

	return true
}

func main() {
	ln1 := insert([]int{1, 2, 2, 1})
	show(ln1)
	fmt.Println(isPalindrome(ln1))
	ln2 := insert([]int{1, 2})
	show(ln2)
	fmt.Println(isPalindrome(ln2))
}
