package main

import (
	"testing"
)

func TestCountNegatives(t *testing.T) {
	type testCase struct {
		Grid           [][]int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Grid:           [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}},
			ExpectedResult: 8,
		},
		{
			Grid:           [][]int{{3, 2}, {1, 0}},
			ExpectedResult: 0,
		},
	}

	for i := range cases {
		got := countNegatives(cases[i].Grid)
		if got != cases[i].ExpectedResult {
			t.Errorf("Grid: %v, expected: %d got: %d", cases[i].Grid, cases[i].ExpectedResult, got)
		}
	}
}
