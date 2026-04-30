package main

func maxAlternatingSum(nums []int) int64 {
	var plus int64 = 0
	var minus int64 = 0

	for i := range nums {
		newPlus := max(plus, minus+int64(nums[i]))
		newMinus := max(minus, plus-int64(nums[i]))

		plus = newPlus
		minus = newMinus
	}

	return plus
}

func main() {

}
