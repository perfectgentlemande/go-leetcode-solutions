package main

import (
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	type testCase struct {
		RowIndex       int
		ExpectedResult []int
	}

	cases := []testCase{
		{
			RowIndex:       3,
			ExpectedResult: []int{1, 3, 3, 1},
		},
		{
			RowIndex:       0,
			ExpectedResult: []int{1},
		},
		{
			RowIndex:       1,
			ExpectedResult: []int{1, 1},
		},
	}

	for i := range cases {
		got := getRow(cases[i].RowIndex)
		if !reflect.DeepEqual(got, cases[i].ExpectedResult) {
			t.Errorf("rowIndex: %d, expected: %v, got: %v", cases[i].RowIndex, cases[i].ExpectedResult, got)
		}
	}

	for i := range cases {
		got := getRowBetter(cases[i].RowIndex)
		if !reflect.DeepEqual(got, cases[i].ExpectedResult) {
			t.Errorf("(better) rowIndex: %d, expected: %v, got: %v", cases[i].RowIndex, cases[i].ExpectedResult, got)
		}
	}
}
