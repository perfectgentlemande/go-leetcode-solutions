package main

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	type testCase struct {
		NumRows        int
		ExpectedResult [][]int
	}

	cases := []testCase{
		{
			NumRows:        5,
			ExpectedResult: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			NumRows:        1,
			ExpectedResult: [][]int{{1}},
		},
	}

	for i := range cases {
		got := generate(cases[i].NumRows)
		if !reflect.DeepEqual(got, cases[i].ExpectedResult) {
			t.Errorf("numRows: %d, expected: %v, got: %v", cases[i].NumRows, cases[i].ExpectedResult, got)
		}
	}
}
