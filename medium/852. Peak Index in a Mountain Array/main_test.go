package main

import (
	"testing"
)

func TestPeakIndexInMountainArray(t *testing.T) {
	type testCase struct {
		Arr            []int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Arr:            []int{0, 1, 0},
			ExpectedResult: 1,
		},
		{
			Arr:            []int{0, 2, 1, 0},
			ExpectedResult: 1,
		},
		{
			Arr:            []int{0, 10, 5, 2},
			ExpectedResult: 1,
		},
		{
			Arr:            []int{0, 3, 5, 9, 10, 2},
			ExpectedResult: 4,
		},
	}

	for i := range cases {
		got := peakIndexInMountainArray(cases[i].Arr)
		if got != cases[i].ExpectedResult {
			t.Errorf("Arr: %v, expected: %d got: %d", cases[i].Arr, cases[i].ExpectedResult, got)
		}
	}
}
