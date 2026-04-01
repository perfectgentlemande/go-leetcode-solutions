package main

import (
	"testing"
)

func TestDuplicateNumbersXOR(t *testing.T) {
	type testCase struct {
		Nums           []int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Nums:           []int{1, 2, 1, 3},
			ExpectedResult: 1,
		},
		{
			Nums:           []int{1, 2, 3},
			ExpectedResult: 0,
		},
		{
			Nums:           []int{1, 2, 2, 1},
			ExpectedResult: 3,
		},
	}

	for i := range cases {
		got := duplicateNumbersXOR(cases[i].Nums)
		if got != cases[i].ExpectedResult {
			t.Errorf("Nums: %v, expected: %v got: %v", cases[i].Nums, cases[i].ExpectedResult, got)
		}
	}
}
