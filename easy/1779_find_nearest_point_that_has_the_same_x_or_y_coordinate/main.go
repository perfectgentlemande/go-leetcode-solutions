package main

import "fmt"

func isValid(x, y int, point []int) bool {
	if x == point[0] || y == point[1] {
		return true
	}

	return false
}
func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}
func manhattanDistance(x, y int, point []int) int {
	return abs(x-point[0]) + abs(y-point[1])
}
func nearestValidPoint(x int, y int, points [][]int) int {
	if len(points) == 0 {
		return -1
	}

	curDistance := -1
	curIndex := -1
	for i := 0; i < len(points); i++ {
		if isValid(x, y, points[i]) {
			newDistance := manhattanDistance(x, y, points[i])
			if curDistance == -1 || newDistance < curDistance {
				curDistance = newDistance
				curIndex = i
			}
		}
	}

	return curIndex
}

func main() {
	fmt.Println(nearestValidPoint(3, 4, [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}}))
	fmt.Println(nearestValidPoint(3, 4, [][]int{{3, 4}}))
	fmt.Println(nearestValidPoint(3, 4, [][]int{{2, 3}}))
}
