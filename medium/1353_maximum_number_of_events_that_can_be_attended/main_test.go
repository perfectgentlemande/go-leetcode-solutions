package main

import (
	"reflect"
	"testing"
)

func TestShipWithinDays(t *testing.T) {
	type testCase struct {
		Events         [][]int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Events:         [][]int{{1, 2}, {2, 3}, {3, 4}},
			ExpectedResult: 3,
		},
		{
			Events:         [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 2}},
			ExpectedResult: 4,
		},
	}

	for i := range cases {
		got := maxEvents(cases[i].Events)
		if !reflect.DeepEqual(got, cases[i].ExpectedResult) {
			t.Errorf("Events: %v, expected: %v got: %v", cases[i].Events, cases[i].ExpectedResult, got)
		}
	}
}
