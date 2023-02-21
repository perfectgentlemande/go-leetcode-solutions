package main

import (
	"testing"
)

func TestSingleNonDuplicate(t *testing.T) {
	type testCase struct {
		Nums           []int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Nums:           []int{1, 1, 2, 3, 3, 4, 4, 8, 8},
			ExpectedResult: 2,
		},
		{
			Nums:           []int{3, 3, 7, 7, 10, 11, 11},
			ExpectedResult: 10,
		},
		{
			Nums:           []int{1},
			ExpectedResult: 1,
		},
		{
			Nums:           []int{1, 1, 2},
			ExpectedResult: 2,
		},
	}

	for i := range cases {
		got := singleNonDuplicate(cases[i].Nums)
		if got != cases[i].ExpectedResult {
			t.Errorf("Nums: %v, expected: %v got: %v", cases[i].Nums, cases[i].ExpectedResult, got)
		}
	}
}
