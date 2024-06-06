package main

import (
	"testing"
)

func TestIsArraySpecial(t *testing.T) {
	type testCase struct {
		Nums           []int
		ExpectedResult bool
	}

	cases := []testCase{
		{
			Nums:           []int{1},
			ExpectedResult: true,
		},
		{
			Nums:           []int{2, 1, 4},
			ExpectedResult: true,
		},
		{
			Nums:           []int{4, 3, 1, 6},
			ExpectedResult: false,
		},
	}

	for i := range cases {
		got := isArraySpecial(cases[i].Nums)
		if got != cases[i].ExpectedResult {
			t.Errorf("Nums: %v, expected: %v got: %v", cases[i].Nums, cases[i].ExpectedResult, got)
		}
	}
}
