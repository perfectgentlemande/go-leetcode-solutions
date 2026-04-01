package main

func duplicateNumbersXOR(nums []int) int {
	res := 0
	numsMap := map[int]int{}

	for _, num := range nums {
		numsMap[num]++
	}

	for num, count := range numsMap {
		if count == 2 {
			res ^= num
		}
	}

	return res
}

func main() {

}
