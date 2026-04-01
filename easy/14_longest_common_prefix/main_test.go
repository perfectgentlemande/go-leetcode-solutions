package main

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	type testCase struct {
		Strs           []string
		ExpectedResult string
	}

	cases := []testCase{
		{
			Strs:           []string{"flower", "flow", "flight"},
			ExpectedResult: "fl",
		},
		{
			Strs:           []string{"dog", "racecar", "car"},
			ExpectedResult: "",
		},
		{
			Strs:           []string{"cir", "car"},
			ExpectedResult: "c",
		},
	}

	for i := range cases {
		got := longestCommonPrefix(cases[i].Strs)
		if got != cases[i].ExpectedResult {
			t.Errorf("Strs: %v, expected: %s got: %s", cases[i].Strs, cases[i].ExpectedResult, got)
		}
	}
}
