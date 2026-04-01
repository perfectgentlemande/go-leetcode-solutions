package main

import "testing"

func TestNearestValidPoint(t *testing.T) {
	type testCase struct {
		X              int
		Y              int
		Points         [][]int
		ExpectedResult int
	}

	cases := []testCase{
		{
			X:              3,
			Y:              4,
			Points:         [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}},
			ExpectedResult: 2,
		},
		{
			X:              3,
			Y:              4,
			Points:         [][]int{{3, 4}},
			ExpectedResult: 0,
		},
		{
			X:              3,
			Y:              4,
			Points:         [][]int{{2, 3}},
			ExpectedResult: -1,
		},
	}

	for i := range cases {
		got := nearestValidPoint(cases[i].X, cases[i].Y, cases[i].Points)
		if got != cases[i].ExpectedResult {
			t.Errorf("X: %d, Y: %d, Points: %v, expected: %d got: %d", cases[i].X, cases[i].Y, cases[i].Points, cases[i].ExpectedResult, got)
		}
	}
}
