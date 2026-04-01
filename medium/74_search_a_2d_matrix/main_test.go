package main

import (
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	type testCase struct {
		Matrix         [][]int
		Target         int
		ExpectedResult bool
	}

	cases := []testCase{
		{
			Matrix:         [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			Target:         3,
			ExpectedResult: true,
		},
		{
			Matrix:         [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			Target:         13,
			ExpectedResult: false,
		},
		{
			Matrix:         [][]int{{1}},
			Target:         2,
			ExpectedResult: false,
		},
	}

	for i := range cases {
		got := searchMatrix(cases[i].Matrix, cases[i].Target)
		if got != cases[i].ExpectedResult {
			t.Errorf("Matrix: %v, target: %d, expected: %t, got: %t", cases[i].Matrix, cases[i].Target, cases[i].ExpectedResult, got)
		}
	}
}
