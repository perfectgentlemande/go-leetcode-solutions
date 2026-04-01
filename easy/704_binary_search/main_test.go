package main

import "testing"

func TestSearch(t *testing.T) {
	type testCase struct {
		Nums           []int
		Target         int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Nums:           []int{-1, 0, 3, 5, 9, 12},
			Target:         9,
			ExpectedResult: 4,
		},
		{
			Nums:           []int{-1, 0, 3, 5, 9, 12},
			Target:         2,
			ExpectedResult: -1,
		},
	}

	for i := range cases {
		got := search(cases[i].Nums, cases[i].Target)
		if got != cases[i].ExpectedResult {
			t.Errorf("Nums: %v, target: %d, expected: %d got: %d", cases[i].Nums, cases[i].Target, cases[i].ExpectedResult, got)
		}
	}
}
