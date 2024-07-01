package main

func sumOfTheDigitsOfHarshadNumber(x int) int {
	var sumOfDigits int

	if x/10 == 0 {
		return x
	}

	res := x
	for res > 0 {
		mod := res % 10
		sumOfDigits += mod
		res /= 10
	}

	if x/sumOfDigits > 0 && x%sumOfDigits == 0 {
		return sumOfDigits
	}

	return -1
}

func main() {

}
