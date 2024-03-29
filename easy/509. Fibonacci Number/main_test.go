package main

import (
	"testing"
)

func TestFibonacciNumber(t *testing.T) {
	type testCase struct {
		N              int
		ExpectedResult int
	}

	cases := []testCase{
		{
			N:              2,
			ExpectedResult: 1,
		},
		{
			N:              3,
			ExpectedResult: 2,
		},
		{
			N:              4,
			ExpectedResult: 3,
		},
	}

	for i := range cases {
		got := fib(cases[i].N)
		if got != cases[i].ExpectedResult {
			t.Errorf("N: %d, expected: %d got: %d", cases[i].N, cases[i].ExpectedResult, got)
		}
	}
	for i := range cases {
		got := fibWithMemoization(cases[i].N)
		if got != cases[i].ExpectedResult {
			t.Errorf("(with memoization) N: %d, expected: %d got: %d", cases[i].N, cases[i].ExpectedResult, got)
		}
	}
	for i := range cases {
		got := fibIterative(cases[i].N)
		if got != cases[i].ExpectedResult {
			t.Errorf("(iterative) N: %d, expected: %d got: %d", cases[i].N, cases[i].ExpectedResult, got)
		}
	}
}
