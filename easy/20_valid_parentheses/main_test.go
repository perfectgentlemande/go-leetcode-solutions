package main

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	type testCase struct {
		S              string
		ExpectedResult bool
	}

	cases := []testCase{
		{
			S:              "()",
			ExpectedResult: true,
		},
		{
			S:              "()[]{}",
			ExpectedResult: true,
		},
		{
			S:              "(]",
			ExpectedResult: false,
		},
		{
			S:              "([)]",
			ExpectedResult: false,
		},
		{
			S:              "{[]}",
			ExpectedResult: true,
		},
	}

	for i := range cases {
		got := isValid(cases[i].S)
		if got != cases[i].ExpectedResult {
			t.Errorf("Nums: %s, expected: %v, got: %v", cases[i].S, cases[i].ExpectedResult, got)
		}
	}
}
