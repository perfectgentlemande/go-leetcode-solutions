package main

import "testing"

func TestFindMin(t *testing.T) {
	type testCase struct {
		Tokens         []string
		ExpectedResult int
	}

	cases := []testCase{
		{
			Tokens:         []string{"2", "1", "+", "3", "*"},
			ExpectedResult: 9,
		},
		{
			Tokens:         []string{"4", "13", "5", "/", "+"},
			ExpectedResult: 6,
		},
		{
			Tokens:         []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			ExpectedResult: 22,
		},
	}

	for i := range cases {
		got := evalRPN(cases[i].Tokens)
		if got != cases[i].ExpectedResult {
			t.Errorf("Tokens: %v,  expected: %d, got: %v", cases[i].Tokens, cases[i].ExpectedResult, got)
		}
	}
}
