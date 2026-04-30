package main

import (
	"reflect"
	"testing"
)

func TestValidIPAddress(t *testing.T) {
	type testCase struct {
		QueryIP        string
		ExpectedResult string
	}

	cases := []testCase{
		{
			QueryIP:        "172.16.254.1",
			ExpectedResult: "IPv4",
		},
		{
			QueryIP:        "2001:0db8:85a3:0:0:8A2E:0370:7334",
			ExpectedResult: "IPv6",
		},
		{
			QueryIP:        "256.256.256.256",
			ExpectedResult: "Neither",
		},
		{
			QueryIP:        "01.01.01.01",
			ExpectedResult: "Neither",
		},
		{
			QueryIP:        "1.0.1.",
			ExpectedResult: "Neither",
		},
		{
			QueryIP:        "2001:db8:85a3:0::8a2E:0370:7334",
			ExpectedResult: "Neither",
		},
	}

	for i := range cases {
		got := validIPAddress(cases[i].QueryIP)
		if !reflect.DeepEqual(got, cases[i].ExpectedResult) {
			t.Errorf("QueryIP: %v, expected: %v got: %v", cases[i].QueryIP, cases[i].ExpectedResult, got)
		}
	}
}
