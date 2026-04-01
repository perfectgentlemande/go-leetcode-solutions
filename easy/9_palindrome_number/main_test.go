package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	type testCase struct {
		N              int
		ExpectedResult bool
	}

	cases := []testCase{
		{
			N:              121,
			ExpectedResult: true,
		},
		{
			N:              -121,
			ExpectedResult: false,
		},
		{
			N:              10,
			ExpectedResult: false,
		},
	}

	for i := range cases {
		got := isPalindrome(cases[i].N)
		if got != cases[i].ExpectedResult {
			t.Errorf("N: %d, expected: %t got: %t", cases[i].N, cases[i].ExpectedResult, got)
		}
	}
}
