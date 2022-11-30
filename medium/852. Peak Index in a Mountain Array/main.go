package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] < arr[mid+1] {
			left = mid + 1
		}
		if arr[mid-1] > arr[mid] {
			right = mid - 1
		}
		if arr[mid-1] < arr[mid] && arr[mid] > arr[mid+1] {
			return mid
		}
	}

	return -1
}

func main() {
	fmt.Println(peakIndexInMountainArray([]int{0, 1, 0}))
	fmt.Println(peakIndexInMountainArray([]int{0, 2, 1, 0}))
	fmt.Println(peakIndexInMountainArray([]int{0, 10, 5, 2}))
}
