package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	numsMap := map[int]struct{}{}
	for _, n := range nums {
		numsMap[n] = struct{}{}
	}

	headWithDummy := &ListNode{
		Next: head,
	}

	cur := headWithDummy

	for cur.Next != nil {
		if _, ToDelete := numsMap[cur.Next.Val]; ToDelete {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return headWithDummy.Next
}

func printList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%d\n", cur.Val)
		cur = cur.Next
	}
}

func main() {
	list := &ListNode{
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
	newList := modifiedList([]int{1, 2, 3}, list)
	printList(newList)

	list = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
						},
					},
				},
			},
		},
	}
	newList = modifiedList([]int{1}, list)
	printList(newList)
}
