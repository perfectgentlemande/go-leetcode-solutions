package main

import (
	"reflect"
	"testing"
)

func TestUniqueXorTriplets(t *testing.T) {
	type testCase struct {
		Nums           []int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Nums:           []int{1, 2},
			ExpectedResult: 2,
		},
		{
			Nums:           []int{3, 1, 2},
			ExpectedResult: 4,
		},
	}

	for i := range cases {
		got := uniqueXorTriplets(cases[i].Nums)
		if !reflect.DeepEqual(got, cases[i].ExpectedResult) {
			t.Errorf("Nums: %v, expected: %v, got: %v", cases[i].Nums, cases[i].ExpectedResult, got)
		}
	}
}
