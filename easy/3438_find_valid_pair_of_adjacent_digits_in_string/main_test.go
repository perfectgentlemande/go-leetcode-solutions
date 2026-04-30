package main

import (
	"testing"
)

func TestFindValidPair(t *testing.T) {
	type testCase struct {
		S              string
		ExpectedResult string
	}

	cases := []testCase{
		{
			S:              "2523533",
			ExpectedResult: "23",
		},
		{
			S:              "221",
			ExpectedResult: "21",
		},
		{
			S:              "22",
			ExpectedResult: "",
		},
	}

	for i := range cases {
		got := findValidPair(cases[i].S)
		if got != cases[i].ExpectedResult {
			t.Errorf("S: %v, expected: %v got: %v", cases[i].S, cases[i].ExpectedResult, got)
		}
	}
}
