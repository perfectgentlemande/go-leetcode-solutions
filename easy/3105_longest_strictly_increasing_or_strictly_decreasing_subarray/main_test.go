package main

import (
	"testing"
)

func TestAddedInteger(t *testing.T) {
	type testCase struct {
		Nums           []int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Nums:           []int{1, 4, 3, 3, 2},
			ExpectedResult: 2,
		},
		{
			Nums:           []int{3, 3, 3, 3},
			ExpectedResult: 1,
		},
		{
			Nums:           []int{3, 2, 1},
			ExpectedResult: 3,
		},
		{
			Nums:           []int{1, 2, 3},
			ExpectedResult: 3,
		},
	}

	for i := range cases {
		got := longestMonotonicSubarray(cases[i].Nums)
		if got != cases[i].ExpectedResult {
			t.Errorf("Nums: %v, expected: %v got: %v", cases[i].Nums, cases[i].ExpectedResult, got)
		}
	}
}
