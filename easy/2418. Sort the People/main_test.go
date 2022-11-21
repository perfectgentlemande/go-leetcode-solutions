package main

import (
	"reflect"
	"testing"
)

func TestSortPeople(t *testing.T) {
	type testCase struct {
		Names          []string
		Heights        []int
		ExpectedResult []string
	}

	cases := []testCase{
		{
			Names:          []string{"Mary", "John", "Emma"},
			Heights:        []int{180, 165, 170},
			ExpectedResult: []string{"Mary", "Emma", "John"},
		},
		{
			Names:          []string{"Alice", "Bob", "Bob"},
			Heights:        []int{155, 185, 150},
			ExpectedResult: []string{"Bob", "Alice", "Bob"},
		},
	}

	for i := range cases {
		copiedNames := make([]string, len(cases[i].Names))
		copiedHeights := make([]int, len(cases[i].Heights))
		copy(copiedNames, cases[i].Names)
		copy(copiedHeights, cases[i].Heights)

		got := sortPeople(copiedNames, copiedHeights)
		if !reflect.DeepEqual(cases[i].ExpectedResult, got) {
			t.Errorf("names: %v heights: %v expected: %v got: %v", cases[i].Names, cases[i].Heights, cases[i].ExpectedResult, got)
		}
	}
}
