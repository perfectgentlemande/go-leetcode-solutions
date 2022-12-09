package main

import "fmt"

func maxDistance(nums1 []int, nums2 []int) int {
	maxDistance := 0
	for i := range nums1 {
		if i > len(nums2) {
			break
		}

		l, r := i, len(nums2)-1

		for l <= r {
			mid := l + (r-l)/2

			if nums2[mid] < nums1[i] {
				r = mid - 1
			} else {
				if mid-i > maxDistance {
					maxDistance = mid - i
				}

				l = mid + 1
			}
		}
	}

	return maxDistance
}

func main() {
	fmt.Println(maxDistance([]int{55, 30, 5, 4, 2}, []int{100, 20, 10, 10, 5}))
	fmt.Println(maxDistance([]int{2, 2, 2}, []int{10, 10, 1}))
	fmt.Println(maxDistance([]int{30, 29, 19, 5}, []int{25, 25, 25, 25, 25}))
}
