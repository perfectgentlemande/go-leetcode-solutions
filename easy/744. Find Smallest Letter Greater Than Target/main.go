package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)-1
	if letters[r] <= target {
		return letters[0]
	}

	curSym := target
	for l <= r {
		mid := l + (r-l)/2

		if letters[mid] <= target {
			l = mid + 1
		}
		if letters[mid] > target {
			curSym = letters[mid]
			r = mid - 1
		}
	}

	return curSym
}

func main() {
	fmt.Println(string(nextGreatestLetter([]byte("cfj"), byte('j'))))
}
