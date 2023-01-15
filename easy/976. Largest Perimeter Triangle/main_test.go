package main

import (
	"testing"
)

func TestLargestPerimeter(t *testing.T) {
	type testCase struct {
		Nums           []int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Nums:           []int{2, 1, 2},
			ExpectedResult: 5,
		},
		{
			Nums:           []int{1, 2, 1, 10},
			ExpectedResult: 0,
		},
	}

	for i := range cases {
		got := largestPerimeter(cases[i].Nums)
		if got != cases[i].ExpectedResult {
			t.Errorf("Nums %v, expected: %d got: %d", cases[i].Nums, cases[i].ExpectedResult, got)
		}
	}
}
