package main

import "fmt"

func subtractProductAndSum(n int) int {
	product := 1
	sum := 0

	for n > 0 {
		digit := n % 10

		product *= digit
		sum += digit

		n = n / 10
	}

	return product - sum
}

func main() {
	fmt.Println(subtractProductAndSum(234))
	fmt.Println(subtractProductAndSum(4421))
}
