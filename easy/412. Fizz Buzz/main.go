package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	res := make([]string, n)

	for i := range res {
		num := i + 1
		switch {
		case num%3 == 0 && num%5 == 0:
			res[i] = "FizzBuzz"
		case num%3 == 0:
			res[i] = "Fizz"
		case num%5 == 0:
			res[i] = "Buzz"
		default:
			res[i] = strconv.Itoa(num)
		}
	}

	return res
}

func main() {
	fmt.Println(fizzBuzz(3))
	fmt.Println(fizzBuzz(5))
	fmt.Println(fizzBuzz(15))
}
