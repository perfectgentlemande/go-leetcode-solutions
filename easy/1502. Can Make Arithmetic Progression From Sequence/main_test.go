package main

import (
	"testing"
)

func TestCanMakeArithmeticProgression(t *testing.T) {
	type testCase struct {
		Arr            []int
		ExpectedResult bool
	}

	cases := []testCase{
		{
			Arr:            []int{3, 5, 1},
			ExpectedResult: true,
		},
		{
			Arr:            []int{1, 2, 4},
			ExpectedResult: false,
		},
	}

	for i := range cases {
		got := canMakeArithmeticProgression(cases[i].Arr)
		if got != cases[i].ExpectedResult {
			t.Errorf("Arr: %v, expected: %t got: %t", cases[i].Arr, cases[i].ExpectedResult, got)
		}
	}
}
