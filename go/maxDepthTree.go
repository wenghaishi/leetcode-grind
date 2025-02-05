/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package main

import (
	"testing"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

 func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	return 1 + max(left, right)
 }

 func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

 func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name:     "Empty tree",
			root:     nil,
			expected: 0,
		},
		{
			name: "Single node",
			root: &TreeNode{Val: 1},
			expected: 1,
		},
		{
			name: "Two levels",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{Val: 2},
			},
			expected: 2,
		},
		{
			name: "Three levels unbalanced",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{Val: 3},
				},
			},
			expected: 3,
		},
		{
			name: "Balanced tree",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: 3,
		},
		{
			name: "Right-skewed tree",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{Val: 4},
					},
				},
			},
			expected: 4,
		},
	}
	

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := maxDepth(test.root)
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}