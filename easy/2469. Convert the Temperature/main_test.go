package main

import (
	"reflect"
	"testing"
)

func TestConvertTemperature(t *testing.T) {
	type testCase struct {
		Celsius        float64
		ExpectedResult []float64
	}

	cases := []testCase{
		{
			Celsius:        36.5,
			ExpectedResult: []float64{309.65, 97.7},
		},
		{
			Celsius:        122.11,
			ExpectedResult: []float64{395.26, 251.798},
		},
	}

	for i := range cases {
		got := convertTemperature(cases[i].Celsius)
		if !reflect.DeepEqual(got, cases[i].ExpectedResult) {
			t.Errorf("Celsius: %v, expected: %v got: %v", cases[i].Celsius, cases[i].ExpectedResult, got)
		}
	}
}
