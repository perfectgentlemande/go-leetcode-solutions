package main

func minimumOperations(nums []int) int {
	res := 0

	for i := range nums {
		if (nums[i]+1)%3 == 0 {
			res++
			continue
		}

		if (nums[i]-1)%3 == 0 {
			res++
			continue
		}
	}

	return res
}

func main() {

}
