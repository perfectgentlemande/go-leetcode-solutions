package main

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	type testCase struct {
		Haystack       string
		Needle         string
		ExpectedResult int
	}

	cases := []testCase{
		{
			Haystack:       "sadbutsad",
			Needle:         "sad",
			ExpectedResult: 0,
		},
		{
			Haystack:       "leetcode",
			Needle:         "leeto",
			ExpectedResult: -1,
		},
	}

	for i := range cases {
		got := strStr(cases[i].Haystack, cases[i].Needle)
		if got != cases[i].ExpectedResult {
			t.Errorf("Haystack: %v, needle: %v, expected: %d, got: %d", cases[i].Haystack, cases[i].Needle, cases[i].ExpectedResult, got)
		}
	}
}
