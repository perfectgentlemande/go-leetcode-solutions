package main

type Frame struct {
	n, num int
}

type Queue []Frame

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}
func (q *Queue) Add(fr Frame) {
	*q = append(*q, fr)
}
func (q *Queue) Remove() (Frame, bool) {
	if q.IsEmpty() {
		return Frame{}, false
	} else {
		element := (*q)[0]
		*q = (*q)[1:]
		return element, true
	}
}

func numSquares(n int) int {
	squares := make([]int, 0, n)

	for i := 1; i*i <= n; i++ {
		squares = append(squares, i*i)
	}

	q := Queue{}
	q.Add(Frame{
		n:   n,
		num: 0,
	})

	visited := map[int]bool{}
	for !q.IsEmpty() {
		top, _ := q.Remove()

		if top.n < 0 || visited[top.n] {
			continue
		} else if top.n == 0 {
			return top.num
		}

		visited[top.n] = true
		for _, sq := range squares {
			q.Add(Frame{
				top.n - sq,
				top.num + 1,
			})
		}
	}

	return -1
}

func main() {

}
