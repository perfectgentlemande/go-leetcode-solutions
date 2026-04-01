package main

import "testing"

func TestMaxDepth(t *testing.T) {
	type testCase struct {
		Root           *Node
		ExpectedResult int
	}

	cases := []testCase{
		{
			Root: &Node{
				Val: 1,
				Children: []*Node{
					&Node{
						Val: 3,
						Children: []*Node{
							&Node{
								Val: 5,
							},
							&Node{
								Val: 6,
							},
						},
					},
					&Node{
						Val: 2,
					},
					&Node{
						Val: 4,
					},
				},
			},
			ExpectedResult: 3,
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
