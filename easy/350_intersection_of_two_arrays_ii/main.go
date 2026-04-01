package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	nums1Map := map[int]int{}
	intersectVals := make([]int, 0)

	for _, n1 := range nums1 {
		nums1Map[n1]++
	}

	for _, n2 := range nums2 {
		if nums1Map[n2] > 0 {
			intersectVals = append(intersectVals, n2)
			nums1Map[n2]--
		}
	}

	return intersectVals
}

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := (right + 1) / 2

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])

	return a
}

func intersectAnother(nums1 []int, nums2 []int) []int {
	nums1 = quickSort(nums1)
	nums2 = quickSort(nums2)

	intersectVals := make([]int, 0)

	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			intersectVals = append(intersectVals, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}

	return intersectVals
}

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	fmt.Println(intersectAnother([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersectAnother([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
