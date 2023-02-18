package main

import "testing"

func TestMinDiffInBST(t *testing.T) {
	type testCase struct {
		Root           *TreeNode
		ExpectedResult int
	}

	cases := []testCase{
		{
			Root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
			ExpectedResult: 1,
		},
		{
			Root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 48,
					Left: &TreeNode{
						Val: 12,
					},
					Right: &TreeNode{
						Val: 49,
					},
				},
			},
			ExpectedResult: 1,
		},
	}

	for i := range cases {
		got := minDiffInBST(cases[i].Root)
		if got != cases[i].ExpectedResult {
			t.Errorf("case: %d, expected: %v, got: %v", i, cases[i].ExpectedResult, got)
		}
	}
}
