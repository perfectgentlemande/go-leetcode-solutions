package main

import "testing"

func TestHasPathSum(t *testing.T) {
	type testCase struct {
		Root           *TreeNode
		TargetSum      int
		ExpectedResult bool
	}

	cases := []testCase{
		{
			Root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 11,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			TargetSum:      22,
			ExpectedResult: true,
		},
		{
			Root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			TargetSum:      5,
			ExpectedResult: false,
		},
		{
			Root:           nil,
			TargetSum:      0,
			ExpectedResult: false,
		},
		{
			Root: &TreeNode{
				Val: 1,
			},
			TargetSum:      1,
			ExpectedResult: true,
		},
		{
			Root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			TargetSum:      1,
			ExpectedResult: false,
		},
	}

	for i := range cases {
		got := hasPathSum(cases[i].Root, cases[i].TargetSum)
		if got != cases[i].ExpectedResult {
			t.Errorf("case: %d, target sum: %d, expected: %v, got: %v", i, cases[i].TargetSum, cases[i].ExpectedResult, got)
		}
	}

	for i := range cases {
		got := hasPathSumIterative(cases[i].Root, cases[i].TargetSum)
		if got != cases[i].ExpectedResult {
			t.Errorf("(iterative) case: %d, target sum: %d, expected: %v, got: %v", i, cases[i].TargetSum, cases[i].ExpectedResult, got)
		}
	}
}
