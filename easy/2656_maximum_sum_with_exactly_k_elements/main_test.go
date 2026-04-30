package main

import (
	"testing"
)

func TestMaximizeSum(t *testing.T) {
	type testCase struct {
		Nums           []int
		K              int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Nums:           []int{1, 2, 3, 4, 5},
			K:              3,
			ExpectedResult: 18,
		},
		{
			Nums:           []int{5, 5, 5},
			K:              2,
			ExpectedResult: 11,
		},
	}

	for i := range cases {
		got := maximizeSum(cases[i].Nums, cases[i].K)
		if got != cases[i].ExpectedResult {
			t.Errorf("nums: %v, k: %v expected: %v got: %v", cases[i].Nums, cases[i].K, cases[i].ExpectedResult, got)
		}
	}
}
