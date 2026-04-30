package main

import (
	"testing"
)

func TestSortPeople(t *testing.T) {
	type testCase struct {
		Logs           [][]int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Logs:           [][]int{{1993, 1999}, {2000, 2010}},
			ExpectedResult: 1993,
		},
		{
			Logs:           [][]int{{1950, 1961}, {1960, 1971}, {1970, 1981}},
			ExpectedResult: 1960,
		},
		{
			Logs:           [][]int{{2033, 2034}, {2039, 2047}, {1998, 2042}, {2047, 2048}, {2025, 2029}, {2005, 2044}, {1990, 1992}, {1952, 1956}, {1984, 2014}},
			ExpectedResult: 2005,
		},
	}

	for i := range cases {
		got := maximumPopulation(cases[i].Logs)
		if got != cases[i].ExpectedResult {
			t.Errorf("logs: %v expected: %v got: %v", cases[i].Logs, cases[i].ExpectedResult, got)
		}
	}
}
