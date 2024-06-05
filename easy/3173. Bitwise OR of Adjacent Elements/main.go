package main

func orArray(nums []int) []int {
	res := make([]int, 0, len(nums)-1)

	for i := 0; i < len(nums)-1; i++ {
		res = append(res, nums[i]|nums[i+1])
	}

	return res
}

func main() {}
