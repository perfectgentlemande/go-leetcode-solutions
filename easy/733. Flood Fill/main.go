package main

type Node struct {
	Val  []int
	Next *Node
}
type Stack struct {
	root *Node
}

func (s *Stack) Push(val []int) {
	nextRoot := s.root

	s.root = &Node{
		Val:  val,
		Next: nextRoot,
	}
}

func (s *Stack) Pop() ([]int, bool) {
	if s.root == nil {
		return nil, false
	}

	curVal := s.root.Val
	s.root = s.root.Next

	return curVal, true
}

func (s *Stack) Top() ([]int, bool) {
	if s.root == nil {
		return nil, false
	}

	return s.root.Val, true
}

func (s *Stack) IsEmpty() bool {
	return s.root == nil
}

// pos x,y
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	height := len(image)
	if height == 0 {
		return image
	}
	width := len(image[0])
	srcColor := image[sr][sc]

	explore := Stack{}
	explore.Push([]int{sr, sc})

	for !explore.IsEmpty() {
		// pop the last element in the stack
		position, _ := explore.Pop()
		i, j := position[0], position[1]

		// top
		if i >= 1 && i-1 <= height-1 && image[i-1][j] == srcColor && image[i-1][j] != color {
			explore.Push([]int{i - 1, j})
		}

		// left
		if j >= 1 && j-1 <= width-1 && image[i][j-1] == srcColor && image[i][j-1] != color {
			explore.Push([]int{i, j - 1})
		}

		// right
		if j+1 <= width-1 && image[i][j+1] == srcColor && image[i][j+1] != color {
			explore.Push([]int{i, j + 1})
		}

		// bottom
		if i+1 <= height-1 && image[i+1][j] == srcColor && image[i+1][j] != color {
			explore.Push([]int{i + 1, j})
		}

		// color element
		image[i][j] = color
	}

	return image
}

func main() {}
