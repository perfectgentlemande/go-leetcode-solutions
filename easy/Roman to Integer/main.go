package main

import "fmt"

func romanToInt(s string) int {
	numbersMap := map[byte]int{
		byte('I'): 1,
		byte('V'): 5,
		byte('X'): 10,
		byte('L'): 50,
		byte('C'): 100,
		byte('D'): 500,
		byte('M'): 1000,
	}

	res := numbersMap[s[len(s)-1]]

	for i := len(s) - 1; i > 0; i-- {
		if numbersMap[s[i]] > numbersMap[s[i-1]] {
			res -= numbersMap[s[i-1]]
		} else {
			res += numbersMap[s[i-1]]
		}
	}

	return res
}

func main() {
	fmt.Println("III")
	fmt.Println("IV")
	fmt.Println("LVIII")
	fmt.Println("MCMXCIV")
}
