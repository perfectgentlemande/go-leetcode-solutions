package main

import (
	"testing"
)

func TestNumberOfPairs(t *testing.T) {
	type testCase struct {
		Num1           []int
		Num2           []int
		K              int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Num1:           []int{1, 3, 4},
			Num2:           []int{1, 3, 4},
			K:              1,
			ExpectedResult: 5,
		},
		{
			Num1:           []int{1, 2, 4, 12},
			Num2:           []int{2, 4},
			K:              3,
			ExpectedResult: 2,
		},
	}

	for i := range cases {
		got := numberOfPairs(cases[i].Num1, cases[i].Num2, cases[i].K)
		if got != cases[i].ExpectedResult {
			t.Errorf("Num1: %v, Num2: %v, K: %v, expected: %v got: %v", cases[i].Num1, cases[i].Num2, cases[i].K, cases[i].ExpectedResult, got)
		}
	}
}
