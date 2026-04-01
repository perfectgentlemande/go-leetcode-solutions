package main

import "testing"

func TestMaxDepth(t *testing.T) {
	type testCase struct {
		Root           *TreeNode
		ExpectedResult int
	}

	cases := []testCase{
		{
			Root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			ExpectedResult: 3,
		},
		{
			Root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			ExpectedResult: 2,
		},
	}

	for i := range cases {
		got := maxDepth(cases[i].Root)
		if got != cases[i].ExpectedResult {
			t.Errorf("case: %d, expected: %d, got: %d", i, cases[i].ExpectedResult, got)
		}
	}
	for i := range cases {
		got := maxDepthIterative(cases[i].Root)
		if got != cases[i].ExpectedResult {
			t.Errorf("(iterative) case: %d, expected: %d, got: %d", i, cases[i].ExpectedResult, got)
		}
	}
}
