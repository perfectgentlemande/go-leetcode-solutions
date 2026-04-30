package main

import (
	"testing"
)

func TestMaxAlternatingSum(t *testing.T) {
	type testCase struct {
		Nums           []int
		ExpectedResult int64
	}

	cases := []testCase{
		{
			Nums:           []int{4, 2, 5, 3},
			ExpectedResult: 7,
		},
		//{
		//	Nums:           []int{5, 6, 7, 8},
		//	ExpectedResult: 8,
		//},
		//{
		//	Nums:           []int{6, 2, 1, 2, 4, 5},
		//	ExpectedResult: 10,
		//},
	}

	for i := range cases {
		got := maxAlternatingSum(cases[i].Nums)
		if got != cases[i].ExpectedResult {
			t.Errorf("Nums: %v, expected: %d got: %d", cases[i].Nums, cases[i].ExpectedResult, got)
		}
	}
}
