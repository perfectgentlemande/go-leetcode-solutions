package main

import (
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	type testCase struct {
		X              float64
		N              int
		ExpectedResult float64
	}

	cases := []testCase{
		{
			X:              2.00000,
			N:              10,
			ExpectedResult: 1024.00000,
		},
		{
			X:              2.10000,
			N:              3,
			ExpectedResult: 9.261000000000001,
		},
		{
			X:              2.00000,
			N:              -2,
			ExpectedResult: 0.25000,
		},
		{
			X:              2.00000,
			N:              10,
			ExpectedResult: 1024.00000,
		},
	}

	for i := range cases {
		got := myPow(cases[i].X, cases[i].N)
		if got != cases[i].ExpectedResult {
			t.Errorf("X: %v, N: %v, expected: %v, got: %v", cases[i].X, cases[i].N, cases[i].ExpectedResult, got)
		}
	}
}
