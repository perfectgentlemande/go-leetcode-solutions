package main

import (
	"testing"
)

func TestSumOfTheDigitsOfHarshadNumber(t *testing.T) {
	type testCase struct {
		X              int
		ExpectedResult int
	}

	cases := []testCase{
		{
			X:              18,
			ExpectedResult: 9,
		},
		{
			X:              23,
			ExpectedResult: -1,
		},
	}

	for i := range cases {
		got := sumOfTheDigitsOfHarshadNumber(cases[i].X)
		if got != cases[i].ExpectedResult {
			t.Errorf("X: %v, expected: %v got: %v", cases[i].X, cases[i].ExpectedResult, got)
		}
	}
}
