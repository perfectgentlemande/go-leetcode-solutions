package main

import (
	"testing"
)

func TestHammingWeight(t *testing.T) {
	type testCase struct {
		Num            uint32
		ExpectedResult int
	}

	cases := []testCase{
		{
			Num:            11,
			ExpectedResult: 3,
		},
		{
			Num:            128,
			ExpectedResult: 1,
		},
		{
			Num:            4294967293,
			ExpectedResult: 31,
		},
	}

	for i := range cases {
		got := hammingWeight(cases[i].Num)
		if got != cases[i].ExpectedResult {
			t.Errorf("Num: %d, expected: %d, got: %d", cases[i].Num, cases[i].ExpectedResult, got)
		}
	}
}
