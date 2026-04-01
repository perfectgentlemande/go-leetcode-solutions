package main

import (
	"reflect"
	"testing"
)

func TestUpdateMatrix(t *testing.T) {
	type testCase struct {
		Mat            [][]int
		ExpectedResult [][]int
	}

	cases := []testCase{
		{
			Mat:            [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			ExpectedResult: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		},
		{
			Mat:            [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}},
			ExpectedResult: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}},
		},
		{
			Mat:            [][]int{{0}, {0}, {0}, {0}, {0}},
			ExpectedResult: [][]int{{0}, {0}, {0}, {0}, {0}},
		},
	}

	for i := range cases {
		got := updateMatrix(cases[i].Mat)
		if !reflect.DeepEqual(got, cases[i].ExpectedResult) {
			t.Errorf("Mat: %v, expected: %v got: %v", cases[i].Mat, cases[i].ExpectedResult, got)
		}
	}
}
