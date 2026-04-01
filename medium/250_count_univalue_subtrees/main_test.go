package main

import "testing"

func TestCountUnivalSubtrees(t *testing.T) {
	type testCase struct {
		Root           *TreeNode
		ExpectedResult int
	}

	cases := []testCase{
		{
			Root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			ExpectedResult: 4,
		},
	}

	for i := range cases {
		got := countUnivalSubtrees(cases[i].Root)
		if got != cases[i].ExpectedResult {
			t.Errorf("case: %d, expected: %d, got: %d", i, cases[i].ExpectedResult, got)
		}
	}
}
