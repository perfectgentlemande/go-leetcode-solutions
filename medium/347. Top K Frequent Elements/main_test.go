package main

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	type testCase struct {
		Nums           []int
		K              int
		ExpectedResult []int
	}

	cases := []testCase{
		{
			Nums:           []int{1, 1, 1, 2, 2, 3},
			K:              2,
			ExpectedResult: []int{1, 2},
		},
		{
			Nums:           []int{1},
			K:              1,
			ExpectedResult: []int{1},
		},
	}

	for i := range cases {
		got := topKFrequent(cases[i].Nums, cases[i].K)
		if !reflect.DeepEqual(got, cases[i].ExpectedResult) {
			t.Errorf("Nums: %v, K: %d, expected: %v got: %d", cases[i].Nums, cases[i].K, cases[i].ExpectedResult, got)
		}
	}
}
