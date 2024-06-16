package main

func modulus(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
func scoreOfString(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return 0
	}

	res := 0

	prev := rune(s[0])
	for i, ch := range s {
		if i == 0 {
			continue
		}

		res += modulus(int(prev - ch))
		prev = ch
	}

	return res
}

func main() {

}
