package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1

	for left <= right && left >= 0 && right <= len(arr)-1 {
		mid := (left + right) / 2

		if len(arr[0:mid]) == 0 {
			if arr[mid+1] < arr[mid] {
				return mid
			} else {
				return mid + 1
			}
		}
		if len(arr[mid:]) == 1 {
			if arr[mid-1] < arr[mid] {
				return mid
			} else {
				return mid - 1
			}
		}

		if arr[mid-1] < arr[mid] && arr[mid] > arr[mid+1] {
			return mid
		}

		if arr[mid-1] < arr[mid] && arr[mid] < arr[mid+1] {
			left = mid + 1
		}
		if arr[mid-1] > arr[mid] && arr[mid] > arr[mid+1] {
			right = mid - 1
		}

	}

	return -1
}

func main() {
	fmt.Println(peakIndexInMountainArray([]int{0, 1, 0}))
	fmt.Println(peakIndexInMountainArray([]int{0, 2, 1, 0}))
	fmt.Println(peakIndexInMountainArray([]int{0, 10, 5, 2}))
	fmt.Println(peakIndexInMountainArray([]int{3, 5, 3, 2, 0}))
}
