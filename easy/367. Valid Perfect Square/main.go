package main

import "fmt"

func isPerfectSquare(num int) bool {
	for i := 0; i*i <= num; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(16)
	fmt.Println(14)
}
