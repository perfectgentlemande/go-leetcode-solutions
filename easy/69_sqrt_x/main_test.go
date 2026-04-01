package main

import (
	"testing"
)

func TestMySqrt(t *testing.T) {
	type testCase struct {
		X              int
		ExpectedResult int
	}

	cases := []testCase{
		{
			X:              4,
			ExpectedResult: 2,
		},
		{
			X:              8,
			ExpectedResult: 2,
		},
		{
			X:              3,
			ExpectedResult: 1,
		},
		{
			X:              5,
			ExpectedResult: 2,
		},
		{
			X:              7,
			ExpectedResult: 2,
		},
		{
			X:              9,
			ExpectedResult: 3,
		},
		{
			X:              10,
			ExpectedResult: 3,
		},
	}

	for i := range cases {
		got := mySqrt(cases[i].X)
		if got != cases[i].ExpectedResult {
			t.Errorf("X: %d, expected: %d, got: %d", cases[i].X, cases[i].ExpectedResult, got)
		}
	}
}
