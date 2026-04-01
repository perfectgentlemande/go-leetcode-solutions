package main

import (
	"testing"
)

func TestCheckIfExist(t *testing.T) {
	type testCase struct {
		Arr            []int
		ExpectedResult bool
	}

	cases := []testCase{
		{
			Arr:            []int{10, 2, 5, 3},
			ExpectedResult: true,
		},
		{
			Arr:            []int{3, 1, 7, 11},
			ExpectedResult: false,
		},
	}

	for i := range cases {
		got := checkIfExist(cases[i].Arr)
		if got != cases[i].ExpectedResult {
			t.Errorf("Arr: %v, expected: %t got: %t", cases[i].Arr, cases[i].ExpectedResult, got)
		}

		got = checkIfExistBetter(cases[i].Arr)
		if got != cases[i].ExpectedResult {
			t.Errorf("(better) Arr: %v, expected: %t got: %t", cases[i].Arr, cases[i].ExpectedResult, got)
		}
	}
}
