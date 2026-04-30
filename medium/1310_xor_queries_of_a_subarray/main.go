package main

func xorQueriesBrute(arr []int, queries [][]int) []int {
	res := make([]int, len(queries))

	for i := range queries {
		left := queries[i][0]
		right := queries[i][1]

		cur := arr[left]
		for j := left + 1; j <= right; j++ {
			cur = cur ^ arr[j]
		}

		res[i] = cur
	}

	return res
}

func xorQueries(arr []int, queries [][]int) []int {
	for i := 1; i < len(arr); i++ {
		arr[i] = arr[i] ^ arr[i-1]
	}

	res := make([]int, len(queries))
	for i := range queries {
		if queries[i][0] == 0 {
			res[i] = arr[queries[i][1]]
		} else {
			res[i] = arr[queries[i][0]-1] ^ arr[queries[i][1]]
		}
	}

	return res
}

func main() {

}
