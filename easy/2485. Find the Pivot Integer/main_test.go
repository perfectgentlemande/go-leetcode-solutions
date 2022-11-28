package main

import (
	"testing"
)

func TestPivotInteger(t *testing.T) {
	type testCase struct {
		N              int
		ExpectedResult int
	}

	cases := []testCase{
		{
			N:              8,
			ExpectedResult: 6,
		},
		{
			N:              1,
			ExpectedResult: 1,
		},
		{
			N:              4,
			ExpectedResult: -1,
		},
	}

	for i := range cases {
		got := pivotInteger(cases[i].N)
		if got != cases[i].ExpectedResult {
			t.Errorf("N: %d, expected: %d got: %d", cases[i].N, cases[i].ExpectedResult, got)
		}
	}
}
