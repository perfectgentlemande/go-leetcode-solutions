package main

import "fmt"

func twoSum(nums []int, target int) []int {
	res := make([]int, 0, 2)

	for i := range nums {
		selectedNum, selectedIndex := nums[i], i

		if selectedNum > target && target > 0 {
			continue
		}
		if selectedNum < target && target < 0 {
			continue
		}

		for j := range nums {
			if j == selectedIndex {
				continue
			}

			if nums[j]+selectedNum == target {
				res = append(res, selectedIndex, j)
				return res
			}
		}
	}

	return res
}

func twoSumFaster(nums []int, target int) []int {
	res := make([]int, 0, 2)
	mappedNums := make(map[int]int, len(nums))

	for i, v := range nums {
		if index, ok := mappedNums[target-v]; ok {
			res = append(res, index, i)
		}

		mappedNums[v] = i
	}

	return res
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
	fmt.Println(twoSum([]int{-1, -2, -3, -4, -5}, -8))

	fmt.Println(twoSumFaster([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSumFaster([]int{3, 2, 4}, 6))
	fmt.Println(twoSumFaster([]int{3, 3}, 6))
	fmt.Println(twoSumFaster([]int{-1, -2, -3, -4, -5}, -8))
}
