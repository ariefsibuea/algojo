package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

func Test_FindTilt(t *testing.T) {
	testcases := []struct {
		input  *leetcode.TreeNode
		output int
	}{
		{
			input:  mockTreeI(),
			output: 1,
		},
		{
			input:  mockTreeII(),
			output: 15,
		},
		{
			input:  mockTreeIII(),
			output: 9,
		},
	}

	soln := leetcode.Solution{}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			out := soln.FindTilt(tc.input)
			require.Equal(t, tc.output, out)
		})
	}
}

func mockTreeI() *leetcode.TreeNode {
	root := &leetcode.TreeNode{Val: 1}
	root.Left = &leetcode.TreeNode{Val: 2}
	root.Right = &leetcode.TreeNode{Val: 3}
	return root
}

func mockTreeII() *leetcode.TreeNode {
	root := &leetcode.TreeNode{Val: 4}
	root.Left = &leetcode.TreeNode{Val: 2}
	root.Left.Left = &leetcode.TreeNode{Val: 3}
	root.Left.Right = &leetcode.TreeNode{Val: 5}
	root.Right = &leetcode.TreeNode{Val: 9}
	root.Right.Right = &leetcode.TreeNode{Val: 7}
	return root
}

func mockTreeIII() *leetcode.TreeNode {
	root := &leetcode.TreeNode{Val: 21}
	root.Left = &leetcode.TreeNode{Val: 7}
	root.Left.Left = &leetcode.TreeNode{Val: 1}
	root.Left.Right = &leetcode.TreeNode{Val: 1}
	root.Left.Left.Left = &leetcode.TreeNode{Val: 3}
	root.Left.Left.Right = &leetcode.TreeNode{Val: 3}
	root.Right = &leetcode.TreeNode{Val: 14}
	root.Right.Left = &leetcode.TreeNode{Val: 2}
	root.Right.Right = &leetcode.TreeNode{Val: 2}
	return root
}
