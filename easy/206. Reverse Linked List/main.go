package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var newHead *ListNode
	cur := head

	for cur != nil {
		temp := cur
		cur = cur.Next
		temp.Next = newHead
		newHead = temp
	}

	return newHead
}

func reverseListAnother(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var newHead *ListNode
	cur := head

	for cur != nil {
		newHead, cur, cur.Next = cur, cur.Next, newHead
	}

	return newHead
}

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func main() {
	items := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	newHead := reverseList(items)
	for newHead != nil {
		fmt.Println(newHead.Val)
		newHead = newHead.Next
	}
}
