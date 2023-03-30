package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	tmp := &ListNode{Next: head}

	curNode := tmp
	nextNode := tmp
	nextNextNode := tmp

	for {
		if curNode == nil {
			break
		}

		t := curNode
		nextNode = curNode.Next

		if nextNode != nil {
			nextNextNode = nextNode.Next
		} else {
			break
		}
		if nextNextNode != nil {
			curNode = nextNextNode.Next
		} else {
			break
		}
		t.Next = nextNextNode
		nextNextNode.Next = nextNode
		nextNode.Next = curNode
		curNode = nextNode
	}

	return tmp.Next
}

func swapPairsRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head, head.Next.Next, head.Next = head.Next, head, head.Next.Next
	head.Next.Next = swapPairsRecursive(head.Next.Next)
	return head
}

func main() {
	root := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}

	swapPairs(root)
}
