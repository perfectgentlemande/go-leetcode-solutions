package main

import "sort"

func addedInteger(nums1 []int, nums2 []int) int {
	if len(nums1) == 0 {
		return -1
	}

	sort.Ints(nums1)
	sort.Ints(nums2)

	return nums2[0] - nums1[0]
}
func main() {

}
