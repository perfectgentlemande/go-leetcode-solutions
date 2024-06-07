package main

func findPermutationDifference(s string, t string) int {
	stmap := map[rune]struct {
		s int
		t int
	}{}

	res := 0

	for i := 0; i < len(s); i++ {
		v := stmap[rune(s[i])]
		v.s = i
		stmap[rune(s[i])] = v

		v = stmap[rune(t[i])]
		v.t = i
		stmap[rune(t[i])] = v
	}

	for _, v := range stmap {
		diff := v.s - v.t
		if diff < 0 {
			diff = -diff
		}

		res += diff
	}

	return res
}
func main() {

}
