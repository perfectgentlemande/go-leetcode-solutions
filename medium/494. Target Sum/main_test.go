package main

import (
	"testing"
)

func TestFindTargetSumWays(t *testing.T) {
	type testCase struct {
		Nums           []int
		Target         int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Nums:           []int{1, 1, 1, 1, 1},
			Target:         3,
			ExpectedResult: 5,
		},
		{
			Nums:           []int{1},
			Target:         1,
			ExpectedResult: 1,
		},
	}

	for i := range cases {
		got := findTargetSumWays(cases[i].Nums, cases[i].Target)
		if got != cases[i].ExpectedResult {
			t.Errorf("Nums: %v, Target: %v, expected: %v got: %v", cases[i].Nums, cases[i].Target, cases[i].ExpectedResult, got)
		}
	}
}
