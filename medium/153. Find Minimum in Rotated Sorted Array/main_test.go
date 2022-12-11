package main

import "testing"

func TestFindMin(t *testing.T) {
	type testCase struct {
		Nums           []int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Nums:           []int{3, 4, 5, 1, 2},
			ExpectedResult: 1,
		},
		{
			Nums:           []int{4, 5, 6, 7, 0, 1, 2},
			ExpectedResult: 0,
		},
		{
			Nums:           []int{11, 13, 15, 17},
			ExpectedResult: 11,
		},
	}

	for i := range cases {
		got := findMin(cases[i].Nums)
		if got != cases[i].ExpectedResult {
			t.Errorf("Nums: %v,  expected: %d, got: %v", cases[i].Nums, cases[i].ExpectedResult, got)
		}
	}
}
