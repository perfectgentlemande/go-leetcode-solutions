package main

import "fmt"

// I did it with quickSort lol
func sortPeople(names []string, heights []int) []string {
	if len(heights) < 2 {
		return names
	}

	left, right := 0, len(heights)-1

	pivot := (right + 1) / 2

	heights[pivot], heights[right] = heights[right], heights[pivot]
	names[pivot], names[right] = names[right], names[pivot]

	for i := range heights {
		if heights[i] > heights[right] {
			heights[left], heights[i] = heights[i], heights[left]
			names[left], names[i] = names[i], names[left]
			left++
		}
	}

	heights[left], heights[right] = heights[right], heights[left]
	names[left], names[right] = names[right], names[left]

	sortPeople(names[:left], heights[:left])
	sortPeople(names[left+1:], heights[left+1:])

	return names
}
func main() {
	fmt.Println(sortPeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170}))
	fmt.Println(sortPeople([]string{"Alice", "Bob", "Bob"}, []int{155, 185, 150}))
}
