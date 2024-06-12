package main

import (
	"testing"
)

func TestAddedInteger(t *testing.T) {
	type testCase struct {
		Word           string
		ExpectedResult int
	}

	cases := []testCase{
		{
			Word:           "aaAbcBC",
			ExpectedResult: 3,
		},
		{
			Word:           "abc",
			ExpectedResult: 0,
		},
		{
			Word:           "abBCab",
			ExpectedResult: 1,
		},
	}

	for i := range cases {
		got := numberOfSpecialChars(cases[i].Word)
		if got != cases[i].ExpectedResult {
			t.Errorf("Word: %v, expected: %v got: %v", cases[i].Word, cases[i].ExpectedResult, got)
		}
	}
}
