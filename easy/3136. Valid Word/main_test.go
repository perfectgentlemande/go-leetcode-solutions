package main

import (
	"testing"
)

func TestSatisfiesConditions(t *testing.T) {
	type testCase struct {
		Word           string
		ExpectedResult bool
	}

	cases := []testCase{
		{
			Word:           "234Adas",
			ExpectedResult: true,
		},
		{
			Word:           "b3",
			ExpectedResult: false,
		},
		{
			Word:           "a3$e",
			ExpectedResult: false,
		},
	}

	for i := range cases {
		got := isValid(cases[i].Word)
		if got != cases[i].ExpectedResult {
			t.Errorf("Word: %v, expected: %v got: %v", cases[i].Word, cases[i].ExpectedResult, got)
		}
	}
}
