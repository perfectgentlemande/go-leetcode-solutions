package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func traverse(node Node, level int, res *[][]int) {
	if len(*res) == level {
		*res = append(*res, make([]int, 0))
	}
	(*res)[level] = append((*res)[level], node.Val)

	for i := range node.Children {
		if node.Children[i] != nil {
			traverse(*node.Children[i], level+1, res)
		}
	}
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	traverse(*root, 0, &res)

	return res
}

type Queue []Node

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}
func (q *Queue) Add(node Node) {
	*q = append([]Node{node}, *q...)
}
func (q *Queue) Remove() (Node, bool) {
	if q.IsEmpty() {
		return Node{}, false
	} else {
		index := len(*q) - 1
		element := (*q)[index]
		*q = (*q)[:index]
		return element, true
	}
}

func levelOrderIterative(root *Node) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	queue := Queue{*root}
	curLevel := 0

	for !queue.IsEmpty() {
		curQueueLen := len(queue)
		res = append(res, make([]int, 0))

		for i := 0; i < curQueueLen; i++ {
			curNode, _ := queue.Remove()
			res[curLevel] = append(res[curLevel], curNode.Val)

			for j := range curNode.Children {
				if curNode.Children[j] != nil {
					queue.Add(*curNode.Children[j])
				}
			}
		}

		curLevel++
	}

	return res
}

type NodeLevel struct {
	Node
	Level int
}

type Deque []NodeLevel

func (d Deque) IsEmpty() bool {
	return len(d) == 0
}
func (d *Deque) PushBack(node NodeLevel) {
	*d = append(*d, node)
}
func (d *Deque) PushFront(node NodeLevel) {
	*d = append([]NodeLevel{node}, *d...)
}
func (d *Deque) PopBack() (NodeLevel, bool) {
	if d.IsEmpty() {
		return NodeLevel{}, false
	} else {
		index := len(*d) - 1
		element := (*d)[index]
		*d = (*d)[:index]
		return element, true
	}
}
func (d *Deque) PopFront() (NodeLevel, bool) {
	if d.IsEmpty() {
		return NodeLevel{}, false
	} else {
		element := (*d)[0]
		*d = (*d)[1:]
		return element, true
	}
}

func levelOrderIterativeOneMore(root *Node) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	deque := Deque{NodeLevel{
		Node:  *root,
		Level: 0,
	}}

	for !deque.IsEmpty() {
		curNodeLevel, _ := deque.PopBack()
		if curNodeLevel.Level == len(res) {
			res = append(res, make([]int, 0))
		}
		res[curNodeLevel.Level] = append(res[curNodeLevel.Level], curNodeLevel.Val)

		for i := range curNodeLevel.Children {
			if curNodeLevel.Children[i] != nil {
				deque.PushFront(NodeLevel{
					Node:  *curNodeLevel.Children[i],
					Level: curNodeLevel.Level + 1,
				})
			}
		}
	}

	return res
}

func main() {
	r := Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{
						Val: 5,
					},
					{
						Val: 6,
					},
				},
			},
			{
				Val: 2,
			},
			{
				Val: 4,
			},
		},
	}
	fmt.Println(levelOrderIterative(&r))

	r = Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{
						Val: 5,
					},
					{
						Val: 6,
					},
				},
			},
			{
				Val: 2,
			},
			{
				Val: 4,
			},
		},
	}
	fmt.Println(levelOrderIterativeOneMore(&r))

	r = Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{
						Val: 5,
					},
					{
						Val: 6,
					},
				},
			},
			{
				Val: 2,
			},
			{
				Val: 4,
			},
		},
	}
	fmt.Println(levelOrder(&r))
}
