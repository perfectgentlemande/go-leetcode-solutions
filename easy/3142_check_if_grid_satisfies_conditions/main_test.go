package main

import (
	"testing"
)

func TestSatisfiesConditions(t *testing.T) {
	type testCase struct {
		Grid           [][]int
		ExpectedResult bool
	}

	cases := []testCase{
		{
			Grid:           [][]int{{1, 0, 2}, {1, 0, 2}},
			ExpectedResult: true,
		},
	}

	for i := range cases {
		got := satisfiesConditions(cases[i].Grid)
		if got != cases[i].ExpectedResult {
			t.Errorf("Grid: %v, expected: %v got: %v", cases[i].Grid, cases[i].ExpectedResult, got)
		}
	}
}
