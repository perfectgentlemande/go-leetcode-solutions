package main

import (
	"reflect"
	"testing"
)

func TestZigzagLevelOrder(t *testing.T) {
	type testCase struct {
		Root           *TreeNode
		ExpectedResult [][]int
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
			ExpectedResult: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			Root: &TreeNode{
				Val: 1,
			},
			ExpectedResult: [][]int{{1}},
		},
		{
			Root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			ExpectedResult: [][]int{{1}, {3, 2}, {4, 5}},
		},
	}

	for i := range cases {
		got := zigzagLevelOrder(cases[i].Root)
		if !reflect.DeepEqual(got, cases[i].ExpectedResult) {
			t.Errorf("Case: %d,  expected: %v, got: %v", i, cases[i].ExpectedResult, got)
		}
	}
}
