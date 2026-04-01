package main

import (
	"testing"
)

func TestAddedInteger(t *testing.T) {
	type testCase struct {
		Grid           [][]byte
		ExpectedResult bool
	}

	cases := []testCase{
		{
			Grid: [][]byte{
				{byte('B'), byte('W'), byte('B')},
				{byte('B'), byte('W'), byte('W')},
				{byte('B'), byte('W'), byte('B')}},
			ExpectedResult: true,
		},
		{
			Grid: [][]byte{
				{byte('B'), byte('W'), byte('B')},
				{byte('W'), byte('B'), byte('W')},
				{byte('B'), byte('W'), byte('B')}},
			ExpectedResult: false,
		},
		{
			Grid: [][]byte{
				{byte('B'), byte('W'), byte('B')},
				{byte('B'), byte('W'), byte('W')},
				{byte('B'), byte('W'), byte('W')}},
			ExpectedResult: true,
		},
		{
			Grid: [][]byte{
				{byte('W'), byte('W'), byte('W')},
				{byte('W'), byte('W'), byte('W')},
				{byte('W'), byte('W'), byte('W')}},
			ExpectedResult: true,
		},
		{
			Grid: [][]byte{
				{byte('B'), byte('B'), byte('B')},
				{byte('B'), byte('B'), byte('B')},
				{byte('B'), byte('B'), byte('B')}},
			ExpectedResult: true,
		},
	}

	for i := range cases {
		got := canMakeSquare(cases[i].Grid)
		if got != cases[i].ExpectedResult {
			t.Errorf("Grid: %v, expected: %v got: %v", cases[i].Grid, cases[i].ExpectedResult, got)
		}
	}

	for i := range cases {
		got := canMakeSquareBetter(cases[i].Grid)
		if got != cases[i].ExpectedResult {
			t.Errorf("(better) Grid: %v, expected: %v got: %v", cases[i].Grid, cases[i].ExpectedResult, got)
		}
	}
}
