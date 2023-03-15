package main

import (
	"reflect"
	"testing"
)

func TestFloodFill(t *testing.T) {
	type testCase struct {
		Image          [][]int
		Sr             int
		Sc             int
		Color          int
		ExpectedResult [][]int
	}

	cases := []testCase{
		{
			Image:          [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
			Sr:             1,
			Sc:             1,
			Color:          2,
			ExpectedResult: [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
		},
		{
			Image:          [][]int{{0, 0, 0}, {0, 0, 0}},
			Sr:             0,
			Sc:             0,
			Color:          0,
			ExpectedResult: [][]int{{0, 0, 0}, {0, 0, 0}},
		},
	}

	for i := range cases {
		got := floodFill(cases[i].Image, cases[i].Sr, cases[i].Sc, cases[i].Color)
		if !reflect.DeepEqual(got, cases[i].ExpectedResult) {
			t.Errorf("Image: %v, sr: %v, sc: %v, color: %v, expected: %v got: %v", cases[i].Image, cases[i].Sr, cases[i].Sc, cases[i].Color, cases[i].ExpectedResult, got)
		}
	}
}
