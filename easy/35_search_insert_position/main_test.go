package main

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	type testCase struct {
		Nums           []int
		Target         int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Nums:           []int{1, 3, 5, 6},
			Target:         5,
			ExpectedResult: 2,
		},
		{
			Nums:           []int{1, 3, 5, 6},
			Target:         2,
			ExpectedResult: 1,
		},
		{
			Nums:           []int{1, 3, 5, 6},
			Target:         7,
			ExpectedResult: 4,
		},
	}

	for i := range cases {
		got := searchInsert(cases[i].Nums, cases[i].Target)
		if got != cases[i].ExpectedResult {
			t.Errorf("Nums: %v, target: %d, expected: %d, got: %d", cases[i].Nums, cases[i].Target, cases[i].ExpectedResult, got)
		}
	}
}
