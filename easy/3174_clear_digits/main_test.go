package main

import (
	"testing"
)

func TestClearDigits(t *testing.T) {
	type testCase struct {
		S              string
		ExpectedResult string
	}

	cases := []testCase{
		{
			S:              "abc",
			ExpectedResult: "abc",
		},
		{
			S:              "cb34",
			ExpectedResult: "",
		},
	}

	for i := range cases {
		got := clearDigits(cases[i].S)
		if got != cases[i].ExpectedResult {
			t.Errorf("S: %v, expected: %v got: %v", cases[i].S, cases[i].ExpectedResult, got)
		}
	}

	for i := range cases {
		got := clearDigitsBetter(cases[i].S)
		if got != cases[i].ExpectedResult {
			t.Errorf("(better) S: %v, expected: %v got: %v", cases[i].S, cases[i].ExpectedResult, got)
		}
	}
}
