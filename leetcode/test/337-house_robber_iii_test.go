package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"algojo.ariefsibuea.dev/leetcode"
)

func Test_Rob(t *testing.T) {
	testcases := []struct {
		input  *leetcode.TreeNode
		output int
	}{
		{
			input:  mockBinaryTreeI(),
			output: 7,
		},
		{
			input:  mockBinaryTreeII(),
			output: 9,
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("case-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			out := leetcode.Rob(tc.input)
			require.Equal(t, tc.output, out)
		})
	}
}

func mockBinaryTreeI() *leetcode.TreeNode {
	binaryTree := leetcode.TreeNode{Val: 3}
	// left branch
	binaryTree.Left = &leetcode.TreeNode{Val: 2}
	binaryTree.Left.Right = &leetcode.TreeNode{Val: 3}
	// right branch
	binaryTree.Right = &leetcode.TreeNode{Val: 3}
	binaryTree.Right.Right = &leetcode.TreeNode{Val: 1}

	return &binaryTree
}

func mockBinaryTreeII() *leetcode.TreeNode {
	binaryTree := leetcode.TreeNode{Val: 3}
	// left branch
	binaryTree.Left = &leetcode.TreeNode{Val: 4}
	binaryTree.Left.Left = &leetcode.TreeNode{Val: 1}
	binaryTree.Left.Right = &leetcode.TreeNode{Val: 3}
	// right branch
	binaryTree.Right = &leetcode.TreeNode{Val: 5}
	binaryTree.Right.Right = &leetcode.TreeNode{Val: 1}

	return &binaryTree
}
