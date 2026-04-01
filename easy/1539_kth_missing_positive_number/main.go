package main

import "fmt"

func findKthPositive(arr []int, k int) int {
	curMissingNum := 0
	curMissingNumIndex := 0

	for i := 1; i < arr[0]; i++ {
		curMissingNum++
		curMissingNumIndex++

		if curMissingNumIndex == k {
			return curMissingNum
		}
	}

	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if diff > 1 {
			curMissingNum = arr[i-1]
		}
		for diff > 1 {
			curMissingNumIndex++
			curMissingNum++

			if curMissingNumIndex == k {
				return curMissingNum
			}

			diff--
		}
	}

	if curMissingNumIndex == k {
		return curMissingNum
	}

	curMissingNum = arr[len(arr)-1]
	for i := 0; i < k-curMissingNumIndex; i++ {
		curMissingNum++
	}

	return curMissingNum
}

func findKthPositiveBinary(arr []int, k int) int {
	l, r := 0, len(arr)

	for l < r {
		mid := (l + r) / 2
		if arr[mid]-mid-1 < k {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return r + k
}

func main() {
	fmt.Println(findKthPositive([]int{2, 3, 4, 7, 11}, 5))
	fmt.Println(findKthPositive([]int{1, 2, 3, 4}, 2))
	fmt.Println(findKthPositive([]int{2}, 1))
	fmt.Println(findKthPositive([]int{3, 10}, 2))
	fmt.Println(findKthPositiveBinary([]int{2, 3, 4, 7, 11}, 5))
	fmt.Println(findKthPositiveBinary([]int{1, 2, 3, 4}, 2))
	fmt.Println(findKthPositiveBinary([]int{2}, 1))
	fmt.Println(findKthPositiveBinary([]int{3, 10}, 2))
}
