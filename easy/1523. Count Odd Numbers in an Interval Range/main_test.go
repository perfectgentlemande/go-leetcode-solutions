package main

import (
	"testing"
)

func TestCountOdds(t *testing.T) {
	type testCase struct {
		Low            int
		High           int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Low:            3,
			High:           7,
			ExpectedResult: 3,
		},
		{
			Low:            8,
			High:           10,
			ExpectedResult: 1,
		},
	}

	for i := range cases {
		got := countOdds(cases[i].Low, cases[i].High)
		if got != cases[i].ExpectedResult {
			t.Errorf("Low: %d, Hight: %d, expected: %d got: %d", cases[i].Low, cases[i].High, cases[i].ExpectedResult, got)
		}
	}
}
