package main

import (
	"reflect"
	"testing"
)

func TestCountUnivalSubtrees(t *testing.T) {
	type testCase struct {
		Rooms          [][]int
		ExpectedResult [][]int
	}

	cases := []testCase{
		{
			Rooms: [][]int{
				{2147483647, -1, 0, 2147483647},
				{2147483647, 2147483647, 2147483647, -1},
				{2147483647, -1, 2147483647, -1},
				{0, -1, 2147483647, 2147483647},
			},
			ExpectedResult: [][]int{
				{3, -1, 0, 1},
				{2, 2, 1, -1},
				{1, -1, 2, -1},
				{0, -1, 3, 4},
			},
		},
		{
			Rooms: [][]int{
				{-1},
			},
			ExpectedResult: [][]int{
				{-1},
			},
		},
	}

	for i := range cases {
		wallsAndGates(cases[i].Rooms)
		if !reflect.DeepEqual(cases[i].Rooms, cases[i].ExpectedResult) {
			t.Errorf("case: %d, expected: %v, got: %v", i, cases[i].ExpectedResult, cases[i].Rooms)
		}
	}
}
