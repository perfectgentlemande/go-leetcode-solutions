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
			Nums:           []int{5, 7, 7, 8, 8, 10},
			Target:         8,
			ExpectedResult: []int{3, 4},
		},
		{
			Nums:           []int{5, 7, 7, 8, 8, 10},
			Target:         6,
			ExpectedResult: []int{-1, -1},
		},
	}

	for i := range cases {
		got := searchRange(cases[i].Nums, cases[i].Target)
		if !reflect.DeepEqual(got, cases[i].ExpectedResult) {
			t.Errorf("Nums: %v, target: %d, expected: %d, got: %v", cases[i].Nums, cases[i].Target, cases[i].ExpectedResult, got)
		}
	}
}
