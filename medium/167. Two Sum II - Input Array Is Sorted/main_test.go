package main

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	type testCase struct {
		Nums           []int
		Target         int
		ExpectedResult []int
	}

	cases := []testCase{
		{
			Nums:           []int{2, 7, 11, 15},
			Target:         9,
			ExpectedResult: []int{1, 2},
		},
		{
			Nums:           []int{2, 3, 4},
			Target:         6,
			ExpectedResult: []int{1, 3},
		},
		{
			Nums:           []int{-1, 0},
			Target:         -1,
			ExpectedResult: []int{1, 2},
		},
		{
			Nums:           []int{0, 0, 3, 4},
			Target:         0,
			ExpectedResult: []int{1, 2},
		},
		{
			Nums:           []int{-1, -1, 1, 1, 1, 1, 1, 1, 1},
			Target:         -2,
			ExpectedResult: []int{1, 2},
		},
	}

	for i := range cases {
		got := twoSum(cases[i].Nums, cases[i].Target)
		if !reflect.DeepEqual(got, cases[i].ExpectedResult) {
			t.Errorf("Nums: %v, target: %d, expected: %d, got: %v", cases[i].Nums, cases[i].Target, cases[i].ExpectedResult, got)
		}
	}
}
