package main

import (
	"testing"
)

func TestClearDigits(t *testing.T) {
	type testCase struct {
		Nums           []int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Nums:           []int{1, 2, 3, 4},
			ExpectedResult: 3,
		},
		{
			Nums:           []int{3, 6, 9},
			ExpectedResult: 0,
		},
	}

	for i := range cases {
		got := minimumOperations(cases[i].Nums)
		if got != cases[i].ExpectedResult {
			t.Errorf("Nums: %v, expected: %v got: %v", cases[i].Nums, cases[i].ExpectedResult, got)
		}
	}
}
