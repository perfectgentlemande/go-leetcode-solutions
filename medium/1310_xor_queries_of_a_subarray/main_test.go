package main

import (
	"reflect"
	"testing"
)

func TestShipWithinDays(t *testing.T) {
	type testCase struct {
		Arr            []int
		Queries        [][]int
		ExpectedResult []int
	}

	cases := []testCase{
		{
			Arr:            []int{1, 3, 4, 8},
			Queries:        [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}},
			ExpectedResult: []int{2, 7, 14, 8},
		},
		{
			Arr:            []int{4, 8, 2, 10},
			Queries:        [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}},
			ExpectedResult: []int{8, 0, 4, 4},
		},
	}

	for i := range cases {
		got := xorQueriesBrute(cases[i].Arr, cases[i].Queries)
		if !reflect.DeepEqual(got, cases[i].ExpectedResult) {
			t.Errorf("(brute) Arr: %v, queries: %v, expected: %v got: %v", cases[i].Arr, cases[i].Queries, cases[i].ExpectedResult, got)
		}

		got = xorQueries(cases[i].Arr, cases[i].Queries)
		if !reflect.DeepEqual(got, cases[i].ExpectedResult) {
			t.Errorf("Arr: %v, queries: %v, expected: %v got: %v", cases[i].Arr, cases[i].Queries, cases[i].ExpectedResult, got)
		}
	}
}
