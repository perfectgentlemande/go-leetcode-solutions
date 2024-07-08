package main

import (
	"testing"
)

func TestClearDigits(t *testing.T) {
	type testCase struct {
		Nums           []int
		ExpectedResult float64
	}

	cases := []testCase{
		{
			Nums:           []int{7, 8, 3, 4, 15, 13, 4, 1},
			ExpectedResult: 5.5,
		},
		{
			Nums:           []int{1, 9, 8, 3, 10, 5},
			ExpectedResult: 5.5,
		},
		{
			Nums:           []int{1, 2, 3, 7, 8, 9},
			ExpectedResult: 5.0,
		},
	}

	for i := range cases {
		got := minimumAverage(cases[i].Nums)
		if got != cases[i].ExpectedResult {
			t.Errorf("Nums: %v, expected: %v got: %v", cases[i].Nums, cases[i].ExpectedResult, got)
		}
	}
}
