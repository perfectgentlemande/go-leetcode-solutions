package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func kSmallestPairsBrute(nums1 []int, nums2 []int, k int) [][]int {
	allPairs := make([][]int, 0)
	for i := range nums1 {
		for j := range nums2 {
			allPairs = append(allPairs, []int{nums1[i], nums2[j]})
		}
	}

	fmt.Println(allPairs)

	sort.Slice(allPairs, func(i, j int) bool {
		return allPairs[i][0]+allPairs[i][1] < allPairs[j][0]+allPairs[j][1]
	})

	return allPairs[:k]
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	result := make([][]int, 0)
	minHeap := &SetHeap{}
	heap.Init(minHeap)

	for i, v1 := range nums1 {
		heap.Push(minHeap, Set{Sum: v1 + nums2[0], I: i, J: 0})
	}

	for !minHeap.Empty() && k > 0 {
		currentMin := heap.Pop(minHeap).(Set)
		i, j := currentMin.I, currentMin.J
		result = append(result, []int{nums1[i], nums2[j]})

		nextElement := j + 1
		if nextElement < len(nums2) {
			heap.Push(minHeap, Set{Sum: nums1[i] + nums2[nextElement], I: i, J: nextElement})
		}

		k--
	}

	return result
}

type Set struct {
	Sum int
	I   int
	J   int
}
type SetHeap []Set

func (h SetHeap) Len() int {
	return len(h)
}

func (h SetHeap) Less(i, j int) bool {
	return h[i].Sum < h[j].Sum
}

func (h SetHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h SetHeap) Empty() bool {
	return len(h) == 0
}

func (h *SetHeap) Push(x interface{}) {
	*h = append(*h, x.(Set))
}
func (h *SetHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {

}
