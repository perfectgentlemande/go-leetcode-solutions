package main

import "fmt"

func countOdds(low int, high int) int {
	counter := 0

	for i := low; i <= high; i++ {
		if i%2 != 0 {
			counter++
		}
	}

	return counter
}

func main() {
	fmt.Println(countOdds(3, 7))
	fmt.Println(countOdds(8, 10))
}
