package main

import (
	"testing"
)

func TestIsPerfectSquare(t *testing.T) {
	type testCase struct {
		Num            int
		ExpectedResult bool
	}

	cases := []testCase{
		{
			Num:            16,
			ExpectedResult: true,
		},
		{
			Num:            14,
			ExpectedResult: false,
		},
	}

	for i := range cases {
		got := isPerfectSquare(cases[i].Num)
		if got != cases[i].ExpectedResult {
			t.Errorf("Num: %d, expected: %t got: %t", cases[i].Num, cases[i].ExpectedResult, got)
		}
	}
	for i := range cases {
		got := isPerfectSquareBinary(cases[i].Num)
		if got != cases[i].ExpectedResult {
			t.Errorf("Num: %d, expected: %t got: %t", cases[i].Num, cases[i].ExpectedResult, got)
		}
	}
}
