package main

import (
	"reflect"
	"testing"
)

func TestOnesMinusZeros(t *testing.T) {
	type testCase struct {
		Grid           [][]int
		ExpectedResult [][]int
	}

	cases := []testCase{
		{
			Grid:           [][]int{{0, 1, 1}, {1, 0, 1}, {0, 0, 1}},
			ExpectedResult: [][]int{{0, 0, 4}, {0, 0, 4}, {-2, -2, 2}},
		},
		{
			Grid:           [][]int{{1, 1, 1}, {1, 1, 1}},
			ExpectedResult: [][]int{{5, 5, 5}, {5, 5, 5}},
		},
	}

	for i := range cases {
		got := onesMinusZeros(cases[i].Grid)
		if !reflect.DeepEqual(got, cases[i].ExpectedResult) {
			t.Errorf("Grid: %v, expected: %v, got: %v", cases[i].Grid, cases[i].ExpectedResult, got)
		}
	}
}

func TestOnesMinusZerosBetter(t *testing.T) {
	type testCase struct {
		Grid           [][]int
		ExpectedResult [][]int
	}

	cases := []testCase{
		{
			Grid:           [][]int{{0, 1, 1}, {1, 0, 1}, {0, 0, 1}},
			ExpectedResult: [][]int{{0, 0, 4}, {0, 0, 4}, {-2, -2, 2}},
		},
		{
			Grid:           [][]int{{1, 1, 1}, {1, 1, 1}},
			ExpectedResult: [][]int{{5, 5, 5}, {5, 5, 5}},
		},
	}

	for i := range cases {
		got := onesMinusZerosBetter(cases[i].Grid)
		if !reflect.DeepEqual(got, cases[i].ExpectedResult) {
			t.Errorf("Grid: %v, expected: %v, got: %v", cases[i].Grid, cases[i].ExpectedResult, got)
		}
	}
}
