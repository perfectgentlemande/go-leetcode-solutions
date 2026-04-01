package main

import (
	"testing"
)

func TestFindPermutationDifference(t *testing.T) {
	type testCase struct {
		S              string
		T              string
		ExpectedResult int
	}

	cases := []testCase{
		{
			S:              "abc",
			T:              "bac",
			ExpectedResult: 2,
		},
		{
			S:              "abcde",
			T:              "edbac",
			ExpectedResult: 12,
		},
	}

	for i := range cases {
		got := findPermutationDifference(cases[i].S, cases[i].T)
		if got != cases[i].ExpectedResult {
			t.Errorf("S: %v, T: %v, expected: %v got: %v", cases[i].S, cases[i].T, cases[i].ExpectedResult, got)
		}
	}
}
