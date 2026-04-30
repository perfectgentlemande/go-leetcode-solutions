package main

import (
	"testing"
)

func TestArrayStringsAreEqual(t *testing.T) {
	type testCase struct {
		Word1          []string
		Word2          []string
		ExpectedResult bool
	}

	cases := []testCase{
		{
			Word1:          []string{"ab", "c"},
			Word2:          []string{"a", "bc"},
			ExpectedResult: true,
		},
		{
			Word1:          []string{"a", "cb"},
			Word2:          []string{"ab", "c"},
			ExpectedResult: false,
		},
		{
			Word1:          []string{"abc", "d", "defg"},
			Word2:          []string{"abcddefg"},
			ExpectedResult: true,
		},
	}

	for i := range cases {
		got := arrayStringsAreEqual(cases[i].Word1, cases[i].Word2)
		if got != cases[i].ExpectedResult {
			t.Errorf("Word1: %v, Word2: %v, expected: %v got: %v", cases[i].Word1, cases[i].Word2, cases[i].ExpectedResult, got)
		}
	}
}
