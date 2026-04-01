package main

import (
	"testing"
)

func TestAddedInteger(t *testing.T) {
	type testCase struct {
		S              string
		ExpectedResult string
	}

	cases := []testCase{
		{
			S:              "1?:?4",
			ExpectedResult: "11:54",
		},
		{
			S:              "0?:5?",
			ExpectedResult: "09:59",
		},
	}

	for i := range cases {
		got := findLatestTime(cases[i].S)
		if got != cases[i].ExpectedResult {
			t.Errorf("S: %v, expected: %v got: %v", cases[i].S, cases[i].ExpectedResult, got)
		}
	}
}
