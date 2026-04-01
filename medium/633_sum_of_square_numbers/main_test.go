package main

import (
	"testing"
)

func TestJudgeSquareSum(t *testing.T) {
	type testCase struct {
		C              int
		ExpectedResult bool
	}

	cases := []testCase{
		{
			C:              5,
			ExpectedResult: true,
		},
		{
			C:              3,
			ExpectedResult: false,
		},
		{
			C:              100,
			ExpectedResult: true,
		},
	}

	for i := range cases {
		got := judgeSquareSum(cases[i].C)
		if got != cases[i].ExpectedResult {
			t.Errorf("C: %v, expected: %v got: %v", cases[i].C, cases[i].ExpectedResult, got)
		}
	}
}
