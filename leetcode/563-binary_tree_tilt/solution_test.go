package binarytreetilt_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "algojo.ariefsibuea.dev/leetcode/563-binary_tree_tilt"
)

func Test_FindTilt(t *testing.T) {
	testcases := []struct {
		input  *lib.TreeNode
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

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i+1), func(t *testing.T) {
			res := lib.FindTilt(testcase.input)
			require.Equal(t, testcase.output, res)
		})
	}
}

func mockTreeI() *lib.TreeNode {
	root := &lib.TreeNode{Val: 1}
	root.Left = &lib.TreeNode{Val: 2}
	root.Right = &lib.TreeNode{Val: 3}
	return root
}

func mockTreeII() *lib.TreeNode {
	root := &lib.TreeNode{Val: 4}
	root.Left = &lib.TreeNode{Val: 2}
	root.Left.Left = &lib.TreeNode{Val: 3}
	root.Left.Right = &lib.TreeNode{Val: 5}
	root.Right = &lib.TreeNode{Val: 9}
	root.Right.Right = &lib.TreeNode{Val: 7}
	return root
}

func mockTreeIII() *lib.TreeNode {
	root := &lib.TreeNode{Val: 21}
	root.Left = &lib.TreeNode{Val: 7}
	root.Left.Left = &lib.TreeNode{Val: 1}
	root.Left.Right = &lib.TreeNode{Val: 1}
	root.Left.Left.Left = &lib.TreeNode{Val: 3}
	root.Left.Left.Right = &lib.TreeNode{Val: 3}
	root.Right = &lib.TreeNode{Val: 14}
	root.Right.Left = &lib.TreeNode{Val: 2}
	root.Right.Right = &lib.TreeNode{Val: 2}
	return root
}
