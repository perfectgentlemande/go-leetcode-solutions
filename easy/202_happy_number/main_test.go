package main

import "testing"

func TestIsHappy(t *testing.T) {
	type testCase struct {
		N              int
		ExpectedResult bool
	}

	cases := []testCase{
		{
			N:              19,
			ExpectedResult: true,
		},
		{
			N:              2,
			ExpectedResult: false,
		},
	}

	for i := range cases {
		got := isHappy(cases[i].N)
		if got != cases[i].ExpectedResult {
			t.Errorf("N: %d, expected: %t, got: %t", cases[i].N, cases[i].ExpectedResult, got)
		}
	}
}
