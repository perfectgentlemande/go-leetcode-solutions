package main

import (
	"testing"
)

func TestFindTheDistanceValue(t *testing.T) {
	type testCase struct {
		Arr1           []int
		Arr2           []int
		D              int
		ExpectedResult int
	}

	cases := []testCase{
		{
			Arr1:           []int{1, 4, 2, 3},
			Arr2:           []int{-4, -3, 6, 10, 20, 30},
			D:              3,
			ExpectedResult: 2,
		},
		{
			Arr1:           []int{2, 1, 100, 3},
			Arr2:           []int{-5, -2, 10, -3, 7},
			D:              6,
			ExpectedResult: 1,
		},
		{
			Arr1:           []int{-803, 715, -224, 909, 121, -296, 872, 807, 715, 407, 94, -8, 572, 90, -520, -867, 485, -918, -827, -728, -653, -659, 865, 102, -564, -452, 554, -320, 229, 36, 722, -478, -247, -307, -304, -767, -404, -519, 776, 933, 236, 596, 954, 464},
			Arr2:           []int{817, 1, -723, 187, 128, 577, -787, -344, -920, -168, -851, -222, 773, 614, -699, 696, -744, -302, -766, 259, 203, 601, 896, -226, -844, 168, 126, -542, 159, -833, 950, -454, -253, 824, -395, 155, 94, 894, -766, -63, 836, -433, -780, 611, -907, 695, -395, -975, 256, 373, -971, -813, -154, -765, 691, 812, 617, -919, -616, -510, 608, 201, -138, -669, -764, -77, -658, 394, -506, -675, 523, 730, -790, -109, 865, 975, -226, 651, 987, 111, 862, 675, -398, 126, -482, 457, -24, -356, -795, -575, 335, -350, -919, -945, -979, 611},
			D:              37,
			ExpectedResult: 0,
		},
		// {
		// 	Arr1:           []int{-3, 10, 2, 8, 0, 10},
		// 	Arr2:           []int{-9, -1, -4, -9, -8},
		// 	D:              9,
		// 	ExpectedResult: 2,
		// },
	}

	for i := range cases {
		got := findTheDistanceValue(cases[i].Arr1, cases[i].Arr2, cases[i].D)
		if got != cases[i].ExpectedResult {
			t.Errorf("Arr1: %v, Arr2: %v, D: %d, expected: %d got: %d", cases[i].Arr1, cases[i].Arr2, cases[i].D, cases[i].ExpectedResult, got)
		}
	}
	for i := range cases {
		got := findTheDistanceValueBinary(cases[i].Arr1, cases[i].Arr2, cases[i].D)
		if got != cases[i].ExpectedResult {
			t.Errorf("(binary) Arr1: %v, Arr2: %v, D: %d, expected: %d got: %d", cases[i].Arr1, cases[i].Arr2, cases[i].D, cases[i].ExpectedResult, got)
		}
	}
}
