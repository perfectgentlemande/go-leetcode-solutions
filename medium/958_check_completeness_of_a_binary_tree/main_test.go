package main

import (
	"testing"
)

func TestIsCompleteTree(t *testing.T) {
	type testCase struct {
		Root           *TreeNode
		ExpectedResult bool
	}

	cases := []testCase{
		{
			Root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
				},
			},
			ExpectedResult: true,
		},
		{
			Root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			ExpectedResult: false,
		},
	}

	for i := range cases {
		got := isCompleteTree(cases[i].Root)
		if got != cases[i].ExpectedResult {
			t.Errorf("Case: %d, expected: %v got: %v", i, cases[i].ExpectedResult, got)
		}
	}
}
