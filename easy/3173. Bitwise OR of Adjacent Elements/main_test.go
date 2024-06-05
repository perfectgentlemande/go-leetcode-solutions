package main

import (
	"reflect"
	"testing"
)

func TestDuplicateNumbersXOR(t *testing.T) {
	type testCase struct {
		Nums           []int
		ExpectedResult []int
	}

	cases := []testCase{
		{
			Nums:           []int{1, 3, 7, 15},
			ExpectedResult: []int{3, 7, 15},
		},
		{
			Nums:           []int{8, 4, 2},
			ExpectedResult: []int{12, 6},
		},
		{
			Nums:           []int{5, 4, 9, 11},
			ExpectedResult: []int{5, 13, 11},
		},
	}

	for i := range cases {
		got := orArray(cases[i].Nums)
		if !reflect.DeepEqual(got, cases[i].ExpectedResult) {
			t.Errorf("Nums: %v, expected: %v got: %v", cases[i].Nums, cases[i].ExpectedResult, got)
		}
	}
}
