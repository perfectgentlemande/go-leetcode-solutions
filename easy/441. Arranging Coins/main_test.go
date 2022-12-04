package main

import (
	"testing"
)

func TestArrangeCoins(t *testing.T) {
	type testCase struct {
		N              int
		ExpectedResult int
	}

	cases := []testCase{
		{
			N:              5,
			ExpectedResult: 2,
		},
		{
			N:              8,
			ExpectedResult: 3,
		},
		{
			N:              1,
			ExpectedResult: 1,
		},
	}

	for i := range cases {
		got := arrangeCoins(cases[i].N)
		if got != cases[i].ExpectedResult {
			t.Errorf("N: %d, expected: %d got: %d", cases[i].N, cases[i].ExpectedResult, got)
		}
	}
	for i := range cases {
		got := arrangeCoinsBinary(cases[i].N)
		if got != cases[i].ExpectedResult {
			t.Errorf("(binsry) N: %d, expected: %d got: %d", cases[i].N, cases[i].ExpectedResult, got)
		}
	}
}
