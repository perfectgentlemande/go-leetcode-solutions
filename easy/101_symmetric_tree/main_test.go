package main

import "testing"

func TestIsSymmetric(t *testing.T) {
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
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 3,
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
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			ExpectedResult: false,
		},
	}

	for i := range cases {
		got := isSymmetric(cases[i].Root)
		if got != cases[i].ExpectedResult {
			t.Errorf("case: %d, expected: %v, got: %v", i, cases[i].ExpectedResult, got)
		}
	}
	for i := range cases {
		got := isSymmetricIterative(cases[i].Root)
		if got != cases[i].ExpectedResult {
			t.Errorf("case: %d, expected: %v, got: %v", i, cases[i].ExpectedResult, got)
		}
	}
}
