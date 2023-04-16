package main

import (
	"testing"
)

func TestKthGrammar(t *testing.T) {
	type testCase struct {
		N              int
		K              int
		ExpectedResult int
	}

	cases := []testCase{
		{
			N:              1,
			K:              1,
			ExpectedResult: 0,
		},
		{
			N:              2,
			K:              1,
			ExpectedResult: 0,
		},
		{
			N:              2,
			K:              2,
			ExpectedResult: 1,
		},
	}

	for i := range cases {
		got := kthGrammar(cases[i].N, cases[i].K)
		if got != cases[i].ExpectedResult {
			t.Errorf("N: %v, K: %v, expected: %v got: %v", cases[i].N, cases[i].K, cases[i].ExpectedResult, got)
		}
	}
}
