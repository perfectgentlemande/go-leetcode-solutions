package main

import (
	"reflect"
	"testing"
)

func TestJudgeSquareSum(t *testing.T) {
	type testCase struct {
		Temperatures   []int
		ExpectedResult []int
	}

	cases := []testCase{
		{
			Temperatures:   []int{73, 74, 75, 71, 69, 72, 76, 73},
			ExpectedResult: []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			Temperatures:   []int{30, 40, 50, 60},
			ExpectedResult: []int{1, 1, 1, 0},
		},
		{
			Temperatures:   []int{30, 60, 90},
			ExpectedResult: []int{1, 1, 0},
		},
	}

	for i := range cases {
		got := dailyTemperatures(cases[i].Temperatures)
		if !reflect.DeepEqual(got, cases[i].ExpectedResult) {
			t.Errorf("Temperatures: %v, expected: %v got: %v", cases[i].Temperatures, cases[i].ExpectedResult, got)
		}
	}
}
