package main

import (
	"testing"
)

func TestAverage(t *testing.T) {
	type testCase struct {
		Path           string
		ExpectedResult bool
	}

	cases := []testCase{
		{
			Path:           "NES",
			ExpectedResult: false,
		},
		{
			Path:           "NESWW",
			ExpectedResult: true,
		},
	}

	for i := range cases {
		got := isPathCrossing(cases[i].Path)
		if got != cases[i].ExpectedResult {
			t.Errorf("Path: %v, expected: %v got: %v", cases[i].Path, cases[i].ExpectedResult, got)
		}
	}
}
