package main

import "fmt"

func checkIfExist(arr []int) bool {
	for i := range arr {
		for j := range arr {
			if i == j {
				continue
			}

			if arr[i] == 2*arr[j] {
				return true
			}
		}
	}

	return false
}

func checkIfExistBetter(arr []int) bool {
	arrMap := map[int]struct{}{}

	for i := range arr {
		_, okM := arrMap[arr[i]*2]
		_, okD := arrMap[arr[i]/2]
		okD = okD && arr[i]%2 == 0

		if okM || okD {
			return true
		}

		arrMap[arr[i]] = struct{}{}
	}

	return false
}

func main() {
	fmt.Println(checkIfExist([]int{10, 2, 5, 3}))
	fmt.Println(checkIfExist([]int{3, 1, 7, 11}))
	fmt.Println(checkIfExistBetter([]int{10, 2, 5, 3}))
	fmt.Println(checkIfExistBetter([]int{3, 1, 7, 11}))
}
