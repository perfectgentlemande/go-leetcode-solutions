package main

import (
	"testing"
)

func TestAverage(t *testing.T) {
	type testCase struct {
		Salary         []int
		ExpectedResult float64
	}

	cases := []testCase{
		{
			Salary:         []int{4000, 3000, 1000, 2000},
			ExpectedResult: 2500.0,
		},
		{
			Salary:         []int{1000, 2000, 3000},
			ExpectedResult: 2000.0,
		},
	}

	for i := range cases {
		got := average(cases[i].Salary)
		if got != cases[i].ExpectedResult {
			t.Errorf("Salary: %v, expected: %v got: %v", cases[i].Salary, cases[i].ExpectedResult, got)
		}
	}
}
