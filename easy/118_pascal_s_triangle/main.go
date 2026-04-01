package main

func generate(numRows int) [][]int {
	res := make([][]int, 0, numRows)

	for i := 0; i < numRows; i++ {
		res = append(res, make([]int, i+1))

		for j := 0; j < i+1; j++ {
			if j == 0 || i == j {
				res[i][j] = 1
			} else {
				res[i][j] = res[i-1][j-1] + res[i-1][j]
			}
		}
	}

	return res
}

func main() {

}
