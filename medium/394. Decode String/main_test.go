package main

import (
	"testing"
)

func TestDecodeString(t *testing.T) {
	type testCase struct {
		S              string
		ExpectedResult string
	}

	cases := []testCase{
		{
			S:              "3[a]2[bc]",
			ExpectedResult: "aaabcbc",
		},
		{
			S:              "3[a2[c]]",
			ExpectedResult: "accaccacc",
		},
		{
			S:              "2[abc]3[cd]ef",
			ExpectedResult: "abcabccdcdcdef",
		},
	}

	for i := range cases {
		got := decodeString(cases[i].S)
		if got != cases[i].ExpectedResult {
			t.Errorf("S: %s, expected: %s got: %s", cases[i].S, cases[i].ExpectedResult, got)
		}
	}
}
