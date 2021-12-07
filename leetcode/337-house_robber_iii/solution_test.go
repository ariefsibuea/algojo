package houserobberiii_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "github.com/ariefsibuea/cobra/leetcode/337-house_robber_iii"
)

func Test_Rob(t *testing.T) {
	testcases := []struct {
		input  *lib.TreeNode
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

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			res := lib.Rob(testcase.input)
			require.Equal(t, testcase.output, res)
		})
	}
}

func mockBinaryTreeI() *lib.TreeNode {
	binaryTree := lib.TreeNode{Val: 3}
	// left branch
	binaryTree.Left = &lib.TreeNode{Val: 2}
	binaryTree.Left.Right = &lib.TreeNode{Val: 3}
	// right branch
	binaryTree.Right = &lib.TreeNode{Val: 3}
	binaryTree.Right.Right = &lib.TreeNode{Val: 1}

	return &binaryTree
}

func mockBinaryTreeII() *lib.TreeNode {
	binaryTree := lib.TreeNode{Val: 3}
	// left branch
	binaryTree.Left = &lib.TreeNode{Val: 4}
	binaryTree.Left.Left = &lib.TreeNode{Val: 1}
	binaryTree.Left.Right = &lib.TreeNode{Val: 3}
	// right branch
	binaryTree.Right = &lib.TreeNode{Val: 5}
	binaryTree.Right.Right = &lib.TreeNode{Val: 1}

	return &binaryTree
}
