package main

import (
	"testing"
)

func TestFindTheDistanceValue(t *testing.T) {
	type testCase struct {
		Arr1           []int
		Arr2           []int
		D              int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Arr1:           []int{1, 4, 2, 3},
			Arr2:           []int{-4, -3, 6, 10, 20, 30},
			D:              3,
			ExpectedResult: 2,
		},
		{
			Arr1:           []int{2, 1, 100, 3},
			Arr2:           []int{-5, -2, 10, -3, 7},
			D:              6,
			ExpectedResult: 1,
		},
		// {
		// 	Arr1:           []int{-3, 10, 2, 8, 0, 10},
		// 	Arr2:           []int{-9, -1, -4, -9, -8},
		// 	D:              9,
		// 	ExpectedResult: 2,
		// },
	}

	for i := range cases {
		got := findTheDistanceValue(cases[i].Arr1, cases[i].Arr2, cases[i].D)
		if got != cases[i].ExpectedResult {
			t.Errorf("Arr1: %v, Arr2: %v, D: %d, expected: %d got: %d", cases[i].Arr1, cases[i].Arr2, cases[i].D, cases[i].ExpectedResult, got)
		}
	}
	for i := range cases {
		got := findTheDistanceValueBinary(cases[i].Arr1, cases[i].Arr2, cases[i].D)
		if got != cases[i].ExpectedResult {
			t.Errorf("(binary) Arr1: %v, Arr2: %v, D: %d, expected: %d got: %d", cases[i].Arr1, cases[i].Arr2, cases[i].D, cases[i].ExpectedResult, got)
		}
	}
}
