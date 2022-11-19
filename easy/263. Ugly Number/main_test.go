package main

import (
	"testing"
)

func TestIsUgly(t *testing.T) {
	type testCase struct {
		N              int
		ExpectedResult bool
	}

	cases := []testCase{
		{
			N:              6,
			ExpectedResult: true,
		},
		{
			N:              1,
			ExpectedResult: true,
		},
		{
			N:              14,
			ExpectedResult: false,
		},
		{
			N:              10,
			ExpectedResult: true,
		},
		{
			N:              13,
			ExpectedResult: false,
		},
		{
			N:              26,
			ExpectedResult: false,
		},
	}

	for i := range cases {
		got := isUgly(cases[i].N)
		if got != cases[i].ExpectedResult {
			t.Errorf("N: %d, expected: %t got: %t", cases[i].N, cases[i].ExpectedResult, got)
		}
	}
}
