package main

import (
	"reflect"
	"testing"
)

type testCase struct {
	Nums           []int
	Target         int
	ExpectedResult []int
}

func getCases() []testCase {
	return []testCase{
		{
			Nums:           []int{2, 7, 11, 15},
			Target:         9,
			ExpectedResult: []int{0, 1},
		},
		{
			Nums:           []int{3, 2, 4},
			Target:         6,
			ExpectedResult: []int{1, 2},
		},
		{
			Nums:           []int{3, 3},
			Target:         6,
			ExpectedResult: []int{0, 1},
		},
		{
			Nums:           []int{-1, -2, -3, -4, -5},
			Target:         -8,
			ExpectedResult: []int{2, 4},
		},
	}
}

func TestTwoSum(t *testing.T) {
	cases := getCases()

	for i := range cases {
		got := twoSum(cases[i].Nums, cases[i].Target)
		if !reflect.DeepEqual(cases[i].ExpectedResult, got) {
			t.Errorf("nums: %v, target: %d expected: %v got: %v", cases[i].Nums, cases[i].Target, cases[i].ExpectedResult, got)
		}
	}
}

func TestTwoSumFaster(t *testing.T) {
	cases := getCases()

	for i := range cases {
		got := twoSumFaster(cases[i].Nums, cases[i].Target)
		if !reflect.DeepEqual(cases[i].ExpectedResult, got) {
			t.Errorf("nums: %v, target: %d expected: %v got: %v", cases[i].Nums, cases[i].Target, cases[i].ExpectedResult, got)
		}
	}
}
