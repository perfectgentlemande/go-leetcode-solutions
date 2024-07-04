package main

func numberOfChild(n int, k int) int {
	i := 0
	asc := true

	for k > 0 {
		k--
		if i == n-1 {
			i--
			asc = false
			continue
		}
		if i == 0 {
			asc = true
		}

		if asc {
			i++
		} else {
			i--
		}
	}

	return i
}

// k = 5 i = 0 i = 1
// k = 4 i = 1 i = 2
// k = 3 i = 2 i = 1
// k = 2 i = 1
// k = 1 i = 0
// k = 0 i = 1

func main() {

}
