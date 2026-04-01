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

func addedIntegerBetter(nums1 []int, nums2 []int) int {
	if len(nums1) == 0 {
		return -1
	}

	smallest1 := nums1[0]
	for _, v := range nums1 {
		if v < smallest1 {
			smallest1 = v
		}
	}

	smallest2 := nums2[0]
	for _, v := range nums2 {
		if v < smallest2 {
			smallest2 = v
		}
	}

	return smallest2 - smallest1
}

func main() {

}
