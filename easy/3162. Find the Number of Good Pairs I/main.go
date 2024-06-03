package main

func numberOfPairs(nums1 []int, nums2 []int, k int) int {
	res := 0

	for i := range nums1 {
		for j := range nums2 {
			if nums1[i]%(nums2[j]*k) == 0 {
				res++
			}
		}
	}

	return res
}

func main() {

}
