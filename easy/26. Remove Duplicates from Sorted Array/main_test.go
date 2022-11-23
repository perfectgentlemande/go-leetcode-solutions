package main

import (
	"testing"
)

type testCase struct {
	Nums           []int
	ExpectedResult int
	ExpectedSlice  []int
}

func getCases() []testCase {
	return []testCase{
		{
			Nums:           []int{1, 1, 2},
			ExpectedResult: 2,
			ExpectedSlice:  []int{1, 2, -1},
		},
		{
			Nums:           []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			ExpectedResult: 5,
			ExpectedSlice:  []int{0, 1, 2, 3, 4, -1, -1, -1, -1, -1},
		},
	}
}

func TestRemoveDuplicates(t *testing.T) {
	cases := getCases()
	for i := range cases {
		copiedNums := make([]int, len(cases[i].Nums))
		copy(copiedNums, cases[i].Nums)

		got := removeDuplicates(copiedNums)
		if got != cases[i].ExpectedResult {
			t.Errorf("Nums: %v, expected: %d, got: %d, expected slice: %v, got: %v", cases[i].Nums, cases[i].ExpectedResult, got, cases[i].ExpectedSlice, copiedNums)
		}
	}
}

func TestRemoveDuplicatesBetter(t *testing.T) {
	cases := getCases()
	for i := range cases {
		copiedNums := make([]int, len(cases[i].Nums))
		copy(copiedNums, cases[i].Nums)

		got := removeDuplicatesBetter(copiedNums)
		if got != cases[i].ExpectedResult {
			t.Errorf("Nums: %v, expected: %d, got: %d, expected slice: %v, got: %v", cases[i].Nums, cases[i].ExpectedResult, got, cases[i].ExpectedSlice, copiedNums)
		}
	}
}
