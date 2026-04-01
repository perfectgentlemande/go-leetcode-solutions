package main

import (
	"testing"
)

func TestGuessNumber(t *testing.T) {
	type testCase struct {
		Pick           int
		N              int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Pick:           6,
			N:              10,
			ExpectedResult: 6,
		},
		{
			Pick:           1,
			N:              1,
			ExpectedResult: 1,
		},
		{
			Pick:           1,
			N:              2,
			ExpectedResult: 1,
		},
	}

	for i := range cases {
		pick = cases[i].Pick
		got := guessNumber(cases[i].N)
		if got != cases[i].ExpectedResult {
			t.Errorf("expected: %d got: %d", cases[i].ExpectedResult, got)
		}
	}
}
