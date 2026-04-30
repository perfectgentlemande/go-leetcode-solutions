package main

func isPathCrossing(path string) bool {
	visited := make(map[[2]int]bool)
	cur := [2]int{0, 0}
	visited[cur] = true

	for _, direction := range path {
		switch direction {
		case 'N':
			cur[1] += 1
		case 'S':
			cur[1] -= 1
		case 'E':
			cur[0] += 1
		case 'W':
			cur[0] -= 1
		}

		if visited[cur] {
			return true
		}

		visited[cur] = true
	}

	return false
}

func main() {

}
