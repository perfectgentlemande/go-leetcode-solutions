package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func rserialize(root *TreeNode, str string) string {
	if root == nil {
		str += "None,"
	} else {
		str += strconv.Itoa(root.Val) + ","
		str = rserialize(root.Left, str)
		str = rserialize(root.Right, str)
	}

	return str
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	return rserialize(root, "")
}

func rdeserialize(l *[]string) *TreeNode {
	if (*l)[0] == "None" {
		*l = (*l)[1:]
		return nil
	}

	v, _ := strconv.Atoi((*l)[0])
	root := &TreeNode{Val: v}

	*l = (*l)[1:]

	root.Left = rdeserialize(l)
	root.Right = rdeserialize(l)

	return root
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	dataList := strings.Split(data, ",")
	return rdeserialize(&dataList)
}

func main() {
	ser := Constructor()
	data := ser.serialize(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	})

	fmt.Println(data)
	root := ser.deserialize(data)
	data = ser.serialize(root)
	fmt.Println(data)
}
