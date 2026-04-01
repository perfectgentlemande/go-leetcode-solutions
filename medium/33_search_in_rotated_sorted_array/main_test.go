package main

import (
	"testing"
)

func TestFindPivot(t *testing.T) {
	type testCase struct {
		Nums           []int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Nums:           []int{4, 5, 6, 7, 0, 1, 2},
			ExpectedResult: 4,
		},
		{
			Nums:           []int{5, 0, 1, 2, 3, 4},
			ExpectedResult: 1,
		},
		{
			Nums:           []int{1},
			ExpectedResult: 0,
		},
		{
			Nums:           []int{5, 1, 3},
			ExpectedResult: 1,
		},
		{
			Nums:           []int{1, 3},
			ExpectedResult: 0,
		},
	}

	for i := range cases {
		got := findPivot(cases[i].Nums)
		if got != cases[i].ExpectedResult {
			t.Errorf("Nums: %v,  expected: %d, got: %v", cases[i].Nums, cases[i].ExpectedResult, got)
		}
	}
}

func TestSearch(t *testing.T) {
	type testCase struct {
		Nums           []int
		Target         int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Nums:           []int{4, 5, 6, 7, 0, 1, 2},
			Target:         0,
			ExpectedResult: 4,
		},
		{
			Nums:           []int{4, 5, 6, 7, 0, 1, 2},
			Target:         3,
			ExpectedResult: -1,
		},
		{
			Nums:           []int{1},
			Target:         0,
			ExpectedResult: -1,
		},
		{
			Nums:           []int{5, 1, 3},
			Target:         1,
			ExpectedResult: 1,
		},
		{
			Nums:           []int{1, 3},
			Target:         3,
			ExpectedResult: 1,
		},
	}

	for i := range cases {
		got := search(cases[i].Nums, cases[i].Target)
		if got != cases[i].ExpectedResult {
			t.Errorf("Nums: %v, target: %d, expected: %d, got: %v", cases[i].Nums, cases[i].Target, cases[i].ExpectedResult, got)
		}
	}
}
