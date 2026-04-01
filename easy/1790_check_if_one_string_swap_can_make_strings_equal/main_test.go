package main

import (
	"testing"
)

func TestAreAlmostEqual(t *testing.T) {
	type testCase struct {
		S1             string
		S2             string
		ExpectedResult bool
	}

	cases := []testCase{
		{
			S1:             "bank",
			S2:             "kanb",
			ExpectedResult: true,
		},
		{
			S1:             "attack",
			S2:             "defend",
			ExpectedResult: false,
		},
		{
			S1:             "kelb",
			S2:             "kelb",
			ExpectedResult: true,
		},
	}

	for i := range cases {
		got := areAlmostEqual(cases[i].S1, cases[i].S2)
		if got != cases[i].ExpectedResult {
			t.Errorf("S1: %s, S2: %s, expected: %t got: %t", cases[i].S1, cases[i].S2, cases[i].ExpectedResult, got)
		}
	}
}
