package main

func uniqueXorTriplets(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}

	res := 1
	for res <= len(nums) {
		res *= 2
	}

	return res
}
func main() {

}
