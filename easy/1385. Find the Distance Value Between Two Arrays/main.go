package main

import (
	"fmt"
)

func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	distance := 0

	for i := range arr1 {
		add := true

		for j := range arr2 {
			if abs(arr1[i]-arr2[j]) <= d {
				add = false
				break
			}
		}

		if add {
			distance++
		}
	}

	return distance
}

func findTheDistanceValueBinary(arr1 []int, arr2 []int, d int) int {
	distance := 0

	for i := range arr1 {
		add := true

		left, right := 0, len(arr2)-1
		for left <= right {
			mid := (left + right) / 2
			fmt.Printf("left: %d, right: %d, mid: %d\n", left, right, mid)

			if abs(arr1[i]-arr2[mid]) <= d {
				add = false
				break
			}
			if abs(arr1[i]-arr2[mid]) > d && arr2[mid] < 0 {
				left = mid + 1
			}
			if abs(arr1[i]-arr2[mid]) > d && arr2[mid] > 0 {
				right = mid - 1
			}
		}
		fmt.Printf("i: %d add: %t\n", i, add)

		if add {
			distance++
		}
	}

	return distance
}

func main() {
	fmt.Println(findTheDistanceValue([]int{4, 5, 8}, []int{10, 9, 1, 8}, 2))
	fmt.Println(findTheDistanceValue([]int{1, 4, 2, 3}, []int{-4, -3, 6, 10, 20, 30}, 3))
	fmt.Println(findTheDistanceValue([]int{2, 1, 100, 3}, []int{-5, -2, 10, -3, 7}, 6))
	fmt.Println(findTheDistanceValue([]int{-3, 10, 2, 8, 0, 10}, []int{-9, -1, -4, -9, -8}, 9))
	fmt.Println(findTheDistanceValueBinary([]int{4, 5, 8}, []int{10, 9, 1, 8}, 2))
	fmt.Println(findTheDistanceValueBinary([]int{1, 4, 2, 3}, []int{-4, -3, 6, 10, 20, 30}, 3))
	fmt.Println(findTheDistanceValueBinary([]int{2, 1, 100, 3}, []int{-5, -2, 10, -3, 7}, 6))
	// this case breaks the binary solution as this is unsorted array
	fmt.Println(findTheDistanceValueBinary([]int{-3, 10, 2, 8, 0, 10}, []int{-9, -1, -4, -9, -8}, 9))
}
