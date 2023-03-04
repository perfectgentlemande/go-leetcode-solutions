package main

import (
	"testing"
)

func TestNumSquares(t *testing.T) {
	type testCase struct {
		N              int
		ExpectedResult int
	}

	cases := []testCase{
		{
			N:              12,
			ExpectedResult: 3,
		},
		{
			N:              13,
			ExpectedResult: 2,
		},
	}

	for i := range cases {
		got := numSquares(cases[i].N)
		if got != cases[i].ExpectedResult {
			t.Errorf("N: %v, expected: %v got: %v", cases[i].N, cases[i].ExpectedResult, got)
		}
	}
}
