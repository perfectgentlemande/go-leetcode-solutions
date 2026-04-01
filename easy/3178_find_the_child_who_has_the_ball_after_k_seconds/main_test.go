package main

import (
	"testing"
)

func TestNumberOfChild(t *testing.T) {
	type testCase struct {
		N              int
		K              int
		ExpectedResult int
	}

	cases := []testCase{
		{
			N:              3,
			K:              5,
			ExpectedResult: 1,
		},
		{
			N:              5,
			K:              6,
			ExpectedResult: 2,
		},
		{
			N:              4,
			K:              2,
			ExpectedResult: 2,
		},
	}

	for i := range cases {
		got := numberOfChild(cases[i].N, cases[i].K)
		if got != cases[i].ExpectedResult {
			t.Errorf("N: %v, K: %v, expected: %v got: %v", cases[i].N, cases[i].K, cases[i].ExpectedResult, got)
		}
	}
}
