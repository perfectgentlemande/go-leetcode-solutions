package main

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	type testCase struct {
		S              []byte
		ExpectedResult []byte
	}

	cases := []testCase{
		{
			S:              []byte("hello"),
			ExpectedResult: []byte("olleh"),
		},
		{
			S:              []byte("Hannah"),
			ExpectedResult: []byte("hannaH"),
		},
	}

	for i := range cases {
		sCopy := make([]byte, len(cases[i].S))
		copy(sCopy, cases[i].S)

		reverseString(sCopy)
		if !reflect.DeepEqual(sCopy, cases[i].ExpectedResult) {
			t.Errorf("s: %v, expected: %v got: %v", cases[i].S, cases[i].ExpectedResult, sCopy)
		}
	}

	for i := range cases {
		sCopy := make([]byte, len(cases[i].S))
		copy(sCopy, cases[i].S)

		reverseStringNonRecursive(sCopy)
		if !reflect.DeepEqual(sCopy, cases[i].ExpectedResult) {
			t.Errorf("(non-recursive) s: %v, expected: %v got: %v", cases[i].S, cases[i].ExpectedResult, sCopy)
		}
	}
}
