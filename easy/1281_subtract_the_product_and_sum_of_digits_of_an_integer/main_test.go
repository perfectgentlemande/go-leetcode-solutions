package main

import (
	"testing"
)

func TestSubtractProductAndSum(t *testing.T) {
	type testCase struct {
		N              int
		ExpectedResult int
	}

	cases := []testCase{
		{
			N:              234,
			ExpectedResult: 15,
		},
		{
			N:              4421,
			ExpectedResult: 21,
		},
	}

	for i := range cases {
		got := subtractProductAndSum(cases[i].N)
		if got != cases[i].ExpectedResult {
			t.Errorf("N: %d, expected: %d got: %d", cases[i].N, cases[i].ExpectedResult, got)
		}
	}
}
