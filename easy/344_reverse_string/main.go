package main

func reverseString(s []byte) {
	if len(s) == 0 {
		return
	}
	if len(s) == 1 {
		return
	}

	s[0], s[len(s)-1] = s[len(s)-1], s[0]
	reverseString(s[1 : len(s)-1])
}

func reverseStringNonRecursive(s []byte) {
	l := 0
	r := len(s) - 1

	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

func main() {

}
