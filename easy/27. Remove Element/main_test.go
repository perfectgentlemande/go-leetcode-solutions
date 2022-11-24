package main

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	type testCase struct {
		Nums           []int
		Val            int
		ExpectedResult int
		ExpectedSlice  []int
	}

	cases := []testCase{
		{
			Nums:           []int{3, 2, 2, 3},
			Val:            3,
			ExpectedResult: 2,
			ExpectedSlice:  []int{2, 2, 2, 3},
		},
		{
			Nums:           []int{0, 1, 2, 2, 3, 0, 4, 2},
			Val:            2,
			ExpectedResult: 5,
			ExpectedSlice:  []int{0, 1, 4, 0, 3, 0, 4, 2},
		},
	}

	for i := range cases {
		copiedNums := make([]int, len(cases[i].Nums))
		copy(copiedNums, cases[i].Nums)

		got := removeElement(copiedNums, cases[i].Val)
		if got != cases[i].ExpectedResult {
			t.Errorf("Nums: %v, val: %d, expected: %d, got: %d, expected slice: %v, got: %v", cases[i].Nums, cases[i].Val, cases[i].ExpectedResult, got, cases[i].ExpectedSlice, copiedNums)
		}
	}
}
