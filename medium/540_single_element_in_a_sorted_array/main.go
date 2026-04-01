package main

func singleNonDuplicate(nums []int) int {
	l := 0
	r := len(nums) - 1

	for l < r {
		m := l + (r-l)/2

		even := (r-m)%2 == 0
		if nums[m] == nums[m+1] {
			if even {
				l = m + 2
			} else {
				r = m - 1
			}
		} else if nums[m-1] == nums[m] {
			if even {
				r = m - 2
			} else {
				l = m + 1
			}
		} else {
			return nums[m]
		}
	}

	return nums[l]
}
func main() {

}
