package main

import (
	"testing"
)

func TestAddedInteger(t *testing.T) {
	type testCase struct {
		Nums1          []int
		Nums2          []int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Nums1:          []int{2, 6, 4},
			Nums2:          []int{9, 7, 5},
			ExpectedResult: 3,
		},
		{
			Nums1:          []int{10},
			Nums2:          []int{5},
			ExpectedResult: -5,
		},
		{
			Nums1:          []int{1, 1, 1, 1},
			Nums2:          []int{1, 1, 1, 1},
			ExpectedResult: 0,
		},
	}

	for i := range cases {
		got := addedInteger(cases[i].Nums1, cases[i].Nums2)
		if got != cases[i].ExpectedResult {
			t.Errorf("nums1: %v, nums2: %v, expected: %v got: %v", cases[i].Nums1, cases[i].Nums2, cases[i].ExpectedResult, got)
		}
	}
}
