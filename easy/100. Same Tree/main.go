package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil && q == nil {
		return false
	}

	if q != nil && p == nil {
		return false
	}

	return (p.Val == q.Val) && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

type Deque []TreeNode

func (d Deque) IsEmpty() bool {
	return len(d) == 0
}
func (d *Deque) PushBack(node TreeNode) {
	*d = append(*d, node)
}
func (d *Deque) PushFront(node TreeNode) {
	*d = append([]TreeNode{node}, *d...)
}
func (d *Deque) PopBack() (TreeNode, bool) {
	if d.IsEmpty() {
		return TreeNode{}, false
	} else {
		index := len(*d) - 1
		element := (*d)[index]
		*d = (*d)[:index]
		return element, true
	}
}
func (d *Deque) PopFront() (TreeNode, bool) {
	if d.IsEmpty() {
		return TreeNode{}, false
	} else {
		element := (*d)[0]
		*d = (*d)[1:]
		return element, true
	}
}

func check(p *TreeNode, q *TreeNode) bool {
	// p and q are null
	if p == nil && q == nil {
		return true
	}
	if q == nil || p == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return true
}

func isSameTreeIterative(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if !check(p, q) {
		return false
	}

	deqP := Deque{}
	deqQ := Deque{}
	deqP.PushBack(*p)
	deqQ.PushBack(*q)

	for !deqP.IsEmpty() {
		pv, _ := deqP.PopFront()
		qv, _ := deqQ.PopFront()

		p = &pv
		q = &qv

		if !check(p, q) {
			return false
		}

		if !check(p.Left, q.Left) {
			return false
		}
		if p.Left != nil {
			deqP.PushBack(*p.Left)
			deqQ.PushBack(*q.Left)
		}

		if !check(p.Right, q.Right) {
			return false
		}

		if p.Right != nil {
			deqP.PushBack(*p.Right)
			deqQ.PushBack(*q.Right)
		}
	}
	return true
}

type Stack []TreeNode

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func (s *Stack) Push(node TreeNode) {
	*s = append(*s, node)
}

func (s *Stack) Pop() (TreeNode, bool) {
	if s.IsEmpty() {
		return TreeNode{}, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func isSameTreeIterativeStack(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if !check(p, q) {
		return false
	}

	sP := Stack{}
	sQ := Stack{}
	sP.Push(*p)
	sQ.Push(*q)

	for !sP.IsEmpty() {
		pv, _ := sP.Pop()
		qv, _ := sQ.Pop()

		p = &pv
		q = &qv

		if !check(p, q) {
			return false
		}

		if !check(p.Right, q.Right) {
			return false
		}

		if p.Right != nil {
			sP.Push(*p.Right)
			sQ.Push(*q.Right)
		}

		if !check(p.Left, q.Left) {
			return false
		}
		if p.Left != nil {
			sP.Push(*p.Left)
			sQ.Push(*q.Left)
		}
	}
	return true
}

func main() {
	p := TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	q := TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(isSameTree(&p, &q))

	p = TreeNode{Val: 1,
		Left: &TreeNode{Val: 2},
	}
	q = TreeNode{Val: 1,
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(isSameTree(&p, &q))

	p = TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	q = TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(isSameTreeIterative(&p, &q))

	p = TreeNode{Val: 1,
		Left: &TreeNode{Val: 2},
	}
	q = TreeNode{Val: 1,
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(isSameTreeIterative(&p, &q))

	p = TreeNode{Val: 1,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
	}
	p = TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(isSameTreeIterative(&p, &q))

	p = TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	q = TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(isSameTreeIterativeStack(&p, &q))

	p = TreeNode{Val: 1,
		Left: &TreeNode{Val: 2},
	}
	q = TreeNode{Val: 1,
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(isSameTreeIterativeStack(&p, &q))

	p = TreeNode{Val: 1,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
	}
	p = TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(isSameTreeIterativeStack(&p, &q))
}
