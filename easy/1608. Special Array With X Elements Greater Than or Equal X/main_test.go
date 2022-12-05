package main

import (
	"testing"
)

func TestSpecialArray(t *testing.T) {
	type testCase struct {
		Nums           []int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Nums:           []int{3, 5},
			ExpectedResult: 2,
		},
		{
			Nums:           []int{0, 0},
			ExpectedResult: -1,
		},
		{
			Nums:           []int{0, 4, 3, 0, 4},
			ExpectedResult: 3,
		},
	}

	for i := range cases {
		got := specialArray(cases[i].Nums)
		if got != cases[i].ExpectedResult {
			t.Errorf("Nums: %v, expected: %d got: %d", cases[i].Nums, cases[i].ExpectedResult, got)
		}
	}
}
