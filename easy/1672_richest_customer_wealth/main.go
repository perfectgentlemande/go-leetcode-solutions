package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	res := 0
	for i := range accounts {
		curRes := 0
		for j := range accounts[i] {
			curRes += accounts[i][j]
		}
		if curRes > res {
			res = curRes
		}
	}

	return res
}

func main() {
	fmt.Println(maximumWealth([][]int{{1, 2, 3}, {3, 2, 1}}))
	fmt.Println(maximumWealth([][]int{{1, 5}, {7, 3}, {3, 5}}))
	fmt.Println(maximumWealth([][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}))
}
