package main

import (
	"testing"
)

func TestMinimumChairs(t *testing.T) {
	type testCase struct {
		S              string
		ExpectedResult int
	}

	cases := []testCase{
		{
			S:              "EEEEEEE",
			ExpectedResult: 7,
		},
		{
			S:              "ELELEEL",
			ExpectedResult: 2,
		},
		{
			S:              "ELEELEELLL",
			ExpectedResult: 3,
		},
		{
			S:              "EELELLEEE",
			ExpectedResult: 3,
		},
	}

	for i := range cases {
		got := minimumChairs(cases[i].S)
		if got != cases[i].ExpectedResult {
			t.Errorf("S: %v, expected: %v got: %v", cases[i].S, cases[i].ExpectedResult, got)
		}
	}
}
