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

// does not affect the initial lists
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	curNode1 := list1
	curNode2 := list2

	var newList *ListNode
	if curNode1 != nil {
		if curNode2 != nil && (curNode2.Val < curNode1.Val) {
			newList = &ListNode{
				Val:  curNode2.Val,
				Next: curNode2.Next,
			}
			curNode2 = curNode2.Next
		} else {
			newList = &ListNode{
				Val:  curNode1.Val,
				Next: curNode1.Next,
			}
			curNode1 = curNode1.Next
		}
	} else if curNode2 != nil {
		if curNode1 != nil && curNode1.Val > curNode2.Val {
			newList = &ListNode{
				Val:  curNode1.Val,
				Next: curNode1.Next,
			}
			curNode1 = curNode1.Next
		} else {
			newList = &ListNode{
				Val:  curNode2.Val,
				Next: curNode2.Next,
			}
			curNode2 = curNode2.Next
		}
	} else {
		return nil
	}

	curNewList := newList
	for curNode1 != nil || curNode2 != nil {
		if curNode1 == nil {
			curNewList.Next = &ListNode{
				Val: curNode2.Val,
			}
			curNewList = curNewList.Next
			curNode2 = curNode2.Next
			continue
		}

		if curNode2 == nil {
			curNewList.Next = &ListNode{
				Val: curNode1.Val,
			}
			curNewList = curNewList.Next
			curNode1 = curNode1.Next
			continue
		}

		if curNode1.Val <= curNode2.Val {
			curNewList.Next = &ListNode{
				Val: curNode1.Val,
			}
			curNewList = curNewList.Next
			curNode1 = curNode1.Next
			continue
		} else {
			curNewList.Next = &ListNode{
				Val: curNode2.Val,
			}
			curNewList = curNewList.Next
			curNode2 = curNode2.Next
			continue
		}
	}

	return newList
}

func mergeTwoListsBetter(list1 *ListNode, list2 *ListNode) *ListNode {
	tempNode := &ListNode{}
	currentNode := tempNode

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			currentNode.Next = list1
			list1 = list1.Next
		} else {
			currentNode.Next = list2
			list2 = list2.Next
		}

		currentNode = currentNode.Next
	}
	if list1 != nil {
		currentNode.Next = list1
		list1 = list1.Next
	}
	if list2 != nil {
		currentNode.Next = list2
		list2 = list2.Next
	}

	return tempNode.Next
}

func main() {
	fmt.Println("unmerged 1:")
	ln1 := insert([]int{1, 2, 4})
	show(ln1)
	fmt.Println("unmerged 2:")
	ln2 := insert([]int{1, 3, 4})
	show(ln2)
	fmt.Println("merged:")
	ln3 := mergeTwoLists(ln1, ln2)
	show(ln3)
	fmt.Println("after merge 1:")
	show(ln1)
	fmt.Println("after merge 2:")
	show(ln2)

	fmt.Println("unmerged 1:")
	ln1 = insert([]int{})
	show(ln1)
	fmt.Println("unmerged 2:")
	ln2 = insert([]int{})
	show(ln2)
	fmt.Println("merged:")
	ln3 = mergeTwoLists(ln1, ln2)
	show(ln3)
	fmt.Println("after merge 1:")
	show(ln1)
	fmt.Println("after merge 2:")
	show(ln2)

	fmt.Println("unmerged 1:")
	ln1 = insert([]int{})
	show(ln1)
	fmt.Println("unmerged 2:")
	ln2 = insert([]int{0})
	show(ln2)
	fmt.Println("merged:")
	ln3 = mergeTwoLists(ln1, ln2)
	show(ln3)
	fmt.Println("after merge 1:")
	show(ln1)
	fmt.Println("after merge 2:")
	show(ln2)

	fmt.Println("unmerged 1:")
	ln1 = insert([]int{2})
	show(ln1)
	fmt.Println("unmerged 2:")
	ln2 = insert([]int{1})
	show(ln2)
	fmt.Println("merged:")
	ln3 = mergeTwoLists(ln1, ln2)
	show(ln3)
	fmt.Println("after merge 1:")
	show(ln1)
	fmt.Println("after merge 2:")
	show(ln2)

	fmt.Println("unmerged 1:")
	ln1 = insert([]int{1, 2, 4})
	show(ln1)
	fmt.Println("unmerged 2:")
	ln2 = insert([]int{1, 3, 4})
	show(ln2)
	fmt.Println("merged:")
	ln3 = mergeTwoListsBetter(ln1, ln2)
	show(ln3)
	fmt.Println("after merge 1:")
	show(ln1)
	fmt.Println("after merge 2:")
	show(ln2)

	fmt.Println("unmerged 1:")
	ln1 = insert([]int{})
	show(ln1)
	fmt.Println("unmerged 2:")
	ln2 = insert([]int{})
	show(ln2)
	fmt.Println("merged:")
	ln3 = mergeTwoListsBetter(ln1, ln2)
	show(ln3)
	fmt.Println("after merge 1:")
	show(ln1)
	fmt.Println("after merge 2:")
	show(ln2)

	fmt.Println("unmerged 1:")
	ln1 = insert([]int{})
	show(ln1)
	fmt.Println("unmerged 2:")
	ln2 = insert([]int{0})
	show(ln2)
	fmt.Println("merged:")
	ln3 = mergeTwoListsBetter(ln1, ln2)
	show(ln3)
	fmt.Println("after merge 1:")
	show(ln1)
	fmt.Println("after merge 2:")
	show(ln2)

	fmt.Println("unmerged 1:")
	ln1 = insert([]int{2})
	show(ln1)
	fmt.Println("unmerged 2:")
	ln2 = insert([]int{1})
	show(ln2)
	fmt.Println("merged:")
	ln3 = mergeTwoListsBetter(ln1, ln2)
	show(ln3)
	fmt.Println("after merge 1:")
	show(ln1)
	fmt.Println("after merge 2:")
	show(ln2)
}
