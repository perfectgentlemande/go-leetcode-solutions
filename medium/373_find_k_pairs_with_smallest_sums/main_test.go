package main

import (
	"reflect"
	"testing"
)

func TestKSmallestPairs(t *testing.T) {
	type testCase struct {
		Nums1          []int
		Nums2          []int
		K              int
		ExpectedResult [][]int
	}

	cases := []testCase{
		{
			Nums1:          []int{1, 7, 11},
			Nums2:          []int{2, 4, 6},
			K:              3,
			ExpectedResult: [][]int{{1, 2}, {1, 4}, {1, 6}},
		},
		{
			Nums1:          []int{1, 1, 2},
			Nums2:          []int{1, 2, 3},
			K:              2,
			ExpectedResult: [][]int{{1, 1}, {1, 1}},
		},
	}

	for i := range cases {
		got := kSmallestPairs(cases[i].Nums1, cases[i].Nums2, cases[i].K)
		if !reflect.DeepEqual(got, cases[i].ExpectedResult) {
			t.Errorf("Nums1: %v, Nums2: %v, K: %d, expected: %v got: %d", cases[i].Nums1, cases[i].Nums2, cases[i].K, cases[i].ExpectedResult, got)
		}
	}
}
