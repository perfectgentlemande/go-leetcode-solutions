package main

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	type testCase struct {
		Prices         []int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Prices:         []int{7, 1, 5, 3, 6, 4},
			ExpectedResult: 5,
		},
		{
			Prices:         []int{7, 6, 4, 3, 1},
			ExpectedResult: 0,
		},
	}

	for i := range cases {
		got := maxProfit(cases[i].Prices)
		if got != cases[i].ExpectedResult {
			t.Errorf("Prices: %v, expected: %d, got: %d", cases[i].Prices, cases[i].ExpectedResult, got)
		}
	}
}
