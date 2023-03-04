package main

import (
	"testing"
)

func TestOpenLock(t *testing.T) {
	type testCase struct {
		Deadends       []string
		Target         string
		ExpectedResult int
	}

	cases := []testCase{
		{
			Deadends:       []string{"0201", "0101", "0102", "1212", "2002"},
			Target:         "0202",
			ExpectedResult: 6,
		},
		{
			Deadends:       []string{"8888"},
			Target:         "0009",
			ExpectedResult: 1,
		},
		{
			Deadends:       []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"},
			Target:         "8888",
			ExpectedResult: -1,
		},
	}

	for i := range cases {
		got := openLock(cases[i].Deadends, cases[i].Target)
		if got != cases[i].ExpectedResult {
			t.Errorf("Deadends: %v, target %v, expected: %v got: %v", cases[i].Deadends, cases[i].Target, cases[i].ExpectedResult, got)
		}
	}
}
