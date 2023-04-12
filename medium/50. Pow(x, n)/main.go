package main

type Pair struct {
	X float64
	N int
}

var myPows map[Pair]float64 = map[Pair]float64{}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n < 0 {
		return 1 / myPow(x, -n)
	}

	pair := Pair{
		X: x,
		N: n,
	}

	res, ok := myPows[pair]
	if !ok {
		if n%2 == 1 {
			res = x * myPow(x, n-1)
		} else {
			res = myPow(x*x, n/2)
		}
		myPows[pair] = res
	}

	return res
}

func main() {

}
