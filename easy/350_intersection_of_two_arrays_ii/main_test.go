package main

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	type testCase struct {
		Nums1          []int
		Nums2          []int
		ExpectedResult []int
	}

	cases := []testCase{
		{
			Nums1:          []int{1, 2, 2, 1},
			Nums2:          []int{2, 2},
			ExpectedResult: []int{2, 2},
		},
		{
			Nums1:          []int{4, 9, 5},
			Nums2:          []int{9, 4, 9, 8, 4},
			ExpectedResult: []int{4, 9},
		},
	}

	for i := range cases {
		got := intersectAnother(cases[i].Nums1, cases[i].Nums2)
		if !reflect.DeepEqual(got, cases[i].ExpectedResult) {
			t.Errorf("Nums1: %v, Nums2: %v, expected: %v got: %v", cases[i].Nums1, cases[i].Nums2, cases[i].ExpectedResult, got)
		}
	}
}
