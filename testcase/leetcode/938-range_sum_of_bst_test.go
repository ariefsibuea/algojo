package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

type inputRangeSumBST struct {
	root *leetcode.TreeNode
	low  int
	high int
}

func Test_RangeSumBST(t *testing.T) {
	testcases := []struct {
		input  inputRangeSumBST
		output int
	}{
		{
			input: inputRangeSumBST{
				root: mockRootRangeSumBSTI(),
				low:  7,
				high: 15,
			},
			output: 32,
		},
		{
			input: inputRangeSumBST{
				root: mockRootRangeSumBSTII(),
				low:  6,
				high: 10,
			},
			output: 23,
		},
	}

	soln := leetcode.Solution{}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i)
		t.Run(testName, func(t *testing.T) {
			require.Equal(t, tc.output, soln.RangeSumBST(tc.input.root, tc.input.low, tc.input.high))
		})
	}
}

func mockRootRangeSumBSTI() *leetcode.TreeNode {
	root := &leetcode.TreeNode{Val: 10}
	root.Left = &leetcode.TreeNode{Val: 5}
	root.Left.Left = &leetcode.TreeNode{Val: 3}
	root.Left.Right = &leetcode.TreeNode{Val: 7}
	root.Right = &leetcode.TreeNode{Val: 15}
	root.Right.Right = &leetcode.TreeNode{Val: 18}
	return root
}

func mockRootRangeSumBSTII() *leetcode.TreeNode {
	root := &leetcode.TreeNode{Val: 10}
	root.Left = &leetcode.TreeNode{Val: 5}
	root.Left.Left = &leetcode.TreeNode{Val: 3}
	root.Left.Right = &leetcode.TreeNode{Val: 7}
	root.Left.Left.Left = &leetcode.TreeNode{Val: 1}
	root.Left.Right.Left = &leetcode.TreeNode{Val: 6}
	root.Right = &leetcode.TreeNode{Val: 15}
	root.Right.Left = &leetcode.TreeNode{Val: 13}
	root.Right.Right = &leetcode.TreeNode{Val: 18}
	return root
}
