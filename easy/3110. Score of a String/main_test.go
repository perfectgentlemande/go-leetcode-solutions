package main

import (
	"testing"
)

func TestAddedInteger(t *testing.T) {
	type testCase struct {
		S              string
		ExpectedResult int
	}

	cases := []testCase{
		{
			S:              "hello",
			ExpectedResult: 13,
		},
		{
			S:              "zaz",
			ExpectedResult: 50,
		},
	}

	for i := range cases {
		got := scoreOfString(cases[i].S)
		if got != cases[i].ExpectedResult {
			t.Errorf("S: %v, expected: %v got: %v", cases[i].S, cases[i].ExpectedResult, got)
		}
	}
}
