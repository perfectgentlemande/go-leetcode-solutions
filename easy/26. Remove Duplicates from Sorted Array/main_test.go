package main

import (
	"reflect"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	type testCase struct {
		Nums                    []int
		ExpectedDistinctNumbers int
		ExpectedResult          []int
	}

	cases := []testCase{
		{
			Nums:                    []int{1, 1, 2},
			ExpectedDistinctNumbers: 2,
			ExpectedResult:          []int{1, 2, -1},
		},
		{
			Nums:                    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			ExpectedDistinctNumbers: 5,
			ExpectedResult:          []int{0, 1, 2, 3, 4, -1, -1, -1, -1, -1},
		},
	}

	for i := range cases {
		copiedNums := make([]int, len(cases[i].Nums))
		copy(copiedNums, cases[i].Nums)

		got := removeDuplicates(copiedNums)
		if !reflect.DeepEqual(got[:cases[i].ExpectedDistinctNumbers], cases[i].ExpectedResult[:cases[i].ExpectedDistinctNumbers]) {
			t.Errorf("Nums: %v, expected: %s got: %s", cases[i].Nums, cases[i].ExpectedResult, got)
		}
	}
}
