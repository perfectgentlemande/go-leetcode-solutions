package main

import (
	"testing"
)

func TestCanVisitAllRooms(t *testing.T) {
	type testCase struct {
		Rooms          [][]int
		ExpectedResult bool
	}

	cases := []testCase{
		{
			Rooms:          [][]int{{1}, {2}, {3}, {}},
			ExpectedResult: true,
		},
		{
			Rooms:          [][]int{{1, 3}, {3, 0, 1}, {2}, {0}},
			ExpectedResult: false,
		},
	}

	for i := range cases {
		got := canVisitAllRooms(cases[i].Rooms)
		if got != cases[i].ExpectedResult {
			t.Errorf("Rooms: %v, expected: %v got: %v", cases[i].Rooms, cases[i].ExpectedResult, got)
		}
	}
}
