package main

import (
	"reflect"
	"testing"
)

func TestKWeakestRows(t *testing.T) {
	type testCase struct {
		Mat            [][]int
		K              int
		ExpectedResult []int
	}

	cases := []testCase{
		{
			Mat:            [][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}},
			K:              3,
			ExpectedResult: []int{2, 0, 3},
		},
		{
			Mat:            [][]int{{1, 0, 0, 0}, {1, 1, 1, 1}, {1, 0, 0, 0}, {1, 0, 0, 0}},
			K:              2,
			ExpectedResult: []int{0, 2},
		},
		{
			Mat:            [][]int{{1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 1, 1, 0}, {0, 0, 0, 0, 0, 0, 0}, {1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 1, 1, 1}},
			K:              4,
			ExpectedResult: []int{2, 0, 3, 1},
		},
	}

	for i := range cases {
		got := kWeakestRows(cases[i].Mat, cases[i].K)
		if !reflect.DeepEqual(got, cases[i].ExpectedResult) {
			t.Errorf("Mat: %v, K: %d, expected: %d got: %d", cases[i].Mat, cases[i].K, cases[i].ExpectedResult, got)
		}
	}

	for i := range cases {
		got := kWeakestRowsBinary(cases[i].Mat, cases[i].K)
		if !reflect.DeepEqual(got, cases[i].ExpectedResult) {
			t.Errorf("(binary) Mat: %v, K: %d, expected: %d got: %d", cases[i].Mat, cases[i].K, cases[i].ExpectedResult, got)
		}
	}
}
