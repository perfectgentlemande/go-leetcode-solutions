package main

import (
	"testing"
)

func TestMaxDistance(t *testing.T) {
	type testCase struct {
		Nums1          []int
		Nums2          []int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Nums1:          []int{55, 30, 5, 4, 2},
			Nums2:          []int{100, 20, 10, 10, 5},
			ExpectedResult: 2,
		},
		{
			Nums1:          []int{2, 2, 2},
			Nums2:          []int{10, 10, 1},
			ExpectedResult: 1,
		},
		{
			Nums1:          []int{30, 29, 19, 5},
			Nums2:          []int{25, 25, 25, 25, 25},
			ExpectedResult: 2,
		},
	}

	for i := range cases {
		got := maxDistance(cases[i].Nums1, cases[i].Nums2)
		if got != cases[i].ExpectedResult {
			t.Errorf("Nums1: %v, Nums2: %v, expected: %d got: %d", cases[i].Nums1, cases[i].Nums2, cases[i].ExpectedResult, got)
		}
	}
}
