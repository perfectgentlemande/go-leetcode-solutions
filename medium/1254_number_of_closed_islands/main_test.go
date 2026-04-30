package main

import (
	"testing"
)

func TestClosedIsland(t *testing.T) {
	type testCase struct {
		Grid           [][]int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Grid: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			ExpectedResult: 1,
		},
		{
			Grid: [][]int{
				{1, 1, 1, 1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0, 1, 1, 0},
				{1, 0, 1, 0, 1, 1, 1, 0},
				{1, 0, 0, 0, 0, 1, 0, 1},
				{1, 1, 1, 1, 1, 1, 1, 0},
			},
			ExpectedResult: 2,
		},
		{
			Grid: [][]int{
				{0, 0, 1, 0, 0},
				{0, 1, 0, 1, 0},
				{0, 1, 1, 1, 0},
			},
			ExpectedResult: 1,
		},
		{
			Grid: [][]int{
				{1, 1, 1, 1, 1, 1, 1},
				{1, 0, 0, 0, 0, 0, 1},
				{1, 0, 1, 1, 1, 0, 1},
				{1, 0, 1, 0, 1, 0, 1},
				{1, 0, 1, 1, 1, 0, 1},
				{1, 0, 0, 0, 0, 0, 1},
				{1, 1, 1, 1, 1, 1, 1},
			},
			ExpectedResult: 2,
		},
	}

	for i := range cases {
		got := closedIsland(cases[i].Grid)
		if got != cases[i].ExpectedResult {
			t.Errorf("Grid: %v, expected: %d got: %d", cases[i].Grid, cases[i].ExpectedResult, got)
		}
	}
}
