package main

import (
	"testing"
)

func TestArraySign(t *testing.T) {
	type testCase struct {
		Nums           []int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Nums:           []int{-1, -2, -3, -4, 3, 2, 1},
			ExpectedResult: 1,
		},
		{
			Nums:           []int{1, 5, 0, 2, -3},
			ExpectedResult: 0,
		},
		{
			Nums:           []int{-1, 1, -1, 1, -1},
			ExpectedResult: -1,
		},
		{
			Nums:           []int{9, 72, 34, 29, -49, -22, -77, -17, -66, -75, -44, -30, -24},
			ExpectedResult: -1,
		},
	}

	for i := range cases {
		got := arraySign(cases[i].Nums)
		if got != cases[i].ExpectedResult {
			t.Errorf("Nums: %v, expected: %d got: %d", cases[i].Nums, cases[i].ExpectedResult, got)
		}
	}
}
