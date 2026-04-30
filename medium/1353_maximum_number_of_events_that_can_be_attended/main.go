package main

import (
	"slices"
)

func maxEvents(events [][]int) int {
	slices.SortFunc(events, func(a, b []int) int {
		return a[0] - b[0]
	})

	return -1
}

func main() {

}
