package main

import "fmt"

func kWeakestRows(mat [][]int, k int) []int {
	rs, is := make([]int, k), make([]int, k)
	for i := range is {
		is[i] = -1
	}

	for i := range mat {
		v := 0
		for v < len(mat[i]) && mat[i][v] > 0 {
			v++
		}

		for j := range rs {
			if is[j] == -1 {
				rs[j] = v
				is[j] = i
				break
			}

			if v < rs[j] || (v == rs[j] && i < is[j]) {
				for k := len(rs) - 1; k > j; k-- {
					rs[k] = rs[k-1]
					is[k] = is[k-1]
				}

				rs[j] = v
				is[j] = i
				break
			}
		}
	}

	return is
}

func kWeakestRowsBinary(mat [][]int, k int) []int {
	rs, is := make([]int, k), make([]int, k)
	for i := range is {
		is[i] = -1
	}

	for i := range mat {
		l, r := 0, len(mat[i])-1

		v := -1
		for l <= r {
			mid := l + (r-l)/2

			if mat[i][mid] > 0 {
				v = mid + 1
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

		for j := range rs {
			if is[j] == -1 {
				rs[j] = v
				is[j] = i
				break
			}

			if v < rs[j] || (v == rs[j] && i < is[j]) {
				for k := len(rs) - 1; k > j; k-- {
					rs[k] = rs[k-1]
					is[k] = is[k-1]
				}

				rs[j] = v
				is[j] = i
				break
			}
		}
	}

	return is
}

func main() {
	fmt.Println(kWeakestRows([][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1}}, 3))
	fmt.Println(kWeakestRows([][]int{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 0, 0, 0},
		{1, 0, 0, 0}}, 2))
	fmt.Println(kWeakestRows([][]int{
		{1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1}}, 4))
	fmt.Println(kWeakestRowsBinary([][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1}}, 3))
	fmt.Println(kWeakestRows([][]int{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 0, 0, 0},
		{1, 0, 0, 0}}, 2))
	fmt.Println(kWeakestRows([][]int{
		{1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1}}, 4))
}
