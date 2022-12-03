package main

import (
	"reflect"
	"testing"
)

func TestNextGreatestLetter(t *testing.T) {
	type testCase struct {
		Letters        []byte
		Target         byte
		ExpectedResult byte
	}

	cases := []testCase{
		{
			Letters:        []byte("cfj"),
			Target:         byte('a'),
			ExpectedResult: byte('c'),
		},
		{
			Letters:        []byte("cfj"),
			Target:         byte('c'),
			ExpectedResult: byte('f'),
		},
		{
			Letters:        []byte("xxyy"),
			Target:         byte('z'),
			ExpectedResult: byte('x'),
		},
		{
			Letters:        []byte("cfj"),
			Target:         byte('d'),
			ExpectedResult: byte('f'),
		},
		{
			Letters:        []byte("cfj"),
			Target:         byte('j'),
			ExpectedResult: byte('c'),
		},
	}

	for i := range cases {
		got := nextGreatestLetter(cases[i].Letters, cases[i].Target)
		if !reflect.DeepEqual(got, cases[i].ExpectedResult) {
			t.Errorf("Letters: %v, target: %v, expected: %v got: %v", cases[i].Letters, cases[i].Target, cases[i].ExpectedResult, got)
		}
	}
}
