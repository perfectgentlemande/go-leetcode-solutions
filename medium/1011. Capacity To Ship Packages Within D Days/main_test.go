package main

import (
	"testing"
)

func TestShipWithinDays(t *testing.T) {
	type testCase struct {
		Arr            []int
		Days           int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Arr:            []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			Days:           5,
			ExpectedResult: 15,
		},
		{
			Arr:            []int{3, 2, 2, 4, 1, 4},
			Days:           3,
			ExpectedResult: 6,
		},
		{
			Arr:            []int{1, 2, 3, 1, 1},
			Days:           4,
			ExpectedResult: 3,
		},
	}

	for i := range cases {
		got := shipWithinDays(cases[i].Arr, cases[i].Days)
		if got != cases[i].ExpectedResult {
			t.Errorf("Arr: %v, days: %d, expected: %d got: %d", cases[i].Arr, cases[i].Days, cases[i].ExpectedResult, got)
		}
	}
}
