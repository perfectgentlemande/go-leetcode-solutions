package main

import (
	"testing"
)

func TestIsUgly(t *testing.T) {
	type testCase struct {
		N              int
		Bad            int
		ExpectedResult int
	}

	cases := []testCase{
		{
			N:              5,
			Bad:            4,
			ExpectedResult: 4,
		},
		{
			N:              1,
			Bad:            1,
			ExpectedResult: 1,
		},
		{
			N:              2147483647,
			Bad:            2147483644,
			ExpectedResult: 2147483644,
		},
	}

	for i := range cases {
		bad = cases[i].Bad

		got := firstBadVersion(cases[i].N)
		if got != cases[i].ExpectedResult {
			t.Errorf("N: %d, Bad: %d, expected: %d got: %d", cases[i].N, cases[i].Bad, cases[i].ExpectedResult, got)
		}
	}
	for i := range cases {
		bad = cases[i].Bad

		got := firstBadVersionBinary(cases[i].N)
		if got != cases[i].ExpectedResult {
			t.Errorf("(binary) N: %d, Bad: %d, expected: %d got: %d", cases[i].N, cases[i].Bad, cases[i].ExpectedResult, got)
		}
	}
}
