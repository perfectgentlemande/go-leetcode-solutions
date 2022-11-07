package main

import "fmt"

func numberOfSteps(num int) int {
	if num == 0 {
		return 0
	}

	steps := 0
	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}

		steps++
	}

	return steps
}

func main() {
	fmt.Println(numberOfSteps(14))
	fmt.Println(numberOfSteps(8))
	fmt.Println(numberOfSteps(123))
}
