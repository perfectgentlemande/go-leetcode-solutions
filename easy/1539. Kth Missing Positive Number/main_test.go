package main

import (
	"testing"
)

func TestFindKthPositive(t *testing.T) {
	type testCase struct {
		Arr            []int
		K              int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Arr:            []int{2, 3, 4, 7, 11},
			K:              5,
			ExpectedResult: 9,
		},
		{
			Arr:            []int{1, 2, 3, 4},
			K:              2,
			ExpectedResult: 6,
		},
		{
			Arr:            []int{2},
			K:              1,
			ExpectedResult: 1,
		},
		{
			Arr:            []int{3, 10},
			K:              2,
			ExpectedResult: 2,
		},
	}

	for i := range cases {
		got := findKthPositive(cases[i].Arr, cases[i].K)
		if got != cases[i].ExpectedResult {
			t.Errorf("Arr: %v, K: %d, expected: %d got: %d", cases[i].Arr, cases[i].K, cases[i].ExpectedResult, got)
		}
	}

	for i := range cases {
		got := findKthPositiveBinary(cases[i].Arr, cases[i].K)
		if got != cases[i].ExpectedResult {
			t.Errorf("(binary) Arr: %v, K: %d, expected: %d got: %d", cases[i].Arr, cases[i].K, cases[i].ExpectedResult, got)
		}
	}
}
