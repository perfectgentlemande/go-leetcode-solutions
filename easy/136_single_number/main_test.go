package main

import (
	"testing"
)

func TestSingleNumber(t *testing.T) {
	type testCase struct {
		Nums           []int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Nums:           []int{2, 2, 1},
			ExpectedResult: 1,
		},
		{
			Nums:           []int{4, 1, 2, 1, 2},
			ExpectedResult: 4,
		},
		{
			Nums:           []int{1},
			ExpectedResult: 1,
		},
	}

	for i := range cases {
		got := singleNumber(cases[i].Nums)
		if got != cases[i].ExpectedResult {
			t.Errorf("Nums: %v, expected: %d, got: %d", cases[i].Nums, cases[i].ExpectedResult, got)
		}
	}
}
