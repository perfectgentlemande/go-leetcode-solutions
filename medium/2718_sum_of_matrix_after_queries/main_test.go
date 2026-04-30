package main

import (
	"testing"
)

func TestOnesMinusZeros(t *testing.T) {
	type testCase struct {
		N              int
		Queries        [][]int
		ExpectedResult int64
	}

	cases := []testCase{
		{
			N:              3,
			Queries:        [][]int{{0, 0, 1}, {1, 2, 2}, {0, 2, 3}, {1, 0, 4}},
			ExpectedResult: 23,
		},
		{
			N:              3,
			Queries:        [][]int{{0, 0, 4}, {0, 1, 2}, {1, 0, 1}, {0, 2, 3}, {1, 2, 1}},
			ExpectedResult: 17,
		},
	}

	for i := range cases {
		got := matrixSumQueries(cases[i].N, cases[i].Queries)
		if got != cases[i].ExpectedResult {
			t.Errorf("N: %v, Queries: %v, expected: %v, got: %v", cases[i].N, cases[i].Queries, cases[i].ExpectedResult, got)
		}
	}
}
