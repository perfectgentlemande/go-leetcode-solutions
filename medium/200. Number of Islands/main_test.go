package main

import (
	"testing"
)

func TestSearchRange(t *testing.T) {
	type testCase struct {
		Grid           [][]byte
		ExpectedResult int
	}

	cases := []testCase{
		{
			Grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			ExpectedResult: 1,
		},
		{
			Grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			ExpectedResult: 3,
		},
	}

	for i := range cases {
		if got := numIslands(cases[i].Grid); got != cases[i].ExpectedResult {
			t.Errorf("Grid %v, expected: %d, got: %v", cases[i].Grid, cases[i].ExpectedResult, got)
		}
	}
}
