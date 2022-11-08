package rangesumofbst_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	lib "algojo.ariefsibuea.dev/leetcode/938-range_sum_of_bst"
)

func Test_RangeSumBST(t *testing.T) {
	testcases := []struct {
		root   *lib.TreeNode
		low    int
		hight  int
		output int
	}{
		{
			root:   mockTreeI(),
			low:    7,
			hight:  15,
			output: 32,
		},
		{
			root:   mockTreeII(),
			low:    6,
			hight:  10,
			output: 23,
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			require.Equal(t, testcase.output, lib.RangeSumBST(testcase.root, testcase.low, testcase.hight))
		})
	}
}

func mockTreeI() *lib.TreeNode {
	root := &lib.TreeNode{Val: 10}
	root.Left = &lib.TreeNode{Val: 5}
	root.Left.Left = &lib.TreeNode{Val: 3}
	root.Left.Right = &lib.TreeNode{Val: 7}
	root.Right = &lib.TreeNode{Val: 15}
	root.Right.Right = &lib.TreeNode{Val: 18}
	return root
}

func mockTreeII() *lib.TreeNode {
	root := &lib.TreeNode{Val: 10}
	root.Left = &lib.TreeNode{Val: 5}
	root.Left.Left = &lib.TreeNode{Val: 3}
	root.Left.Right = &lib.TreeNode{Val: 7}
	root.Left.Left.Left = &lib.TreeNode{Val: 1}
	root.Left.Right.Left = &lib.TreeNode{Val: 6}
	root.Right = &lib.TreeNode{Val: 15}
	root.Right.Left = &lib.TreeNode{Val: 13}
	root.Right.Right = &lib.TreeNode{Val: 18}
	return root
}
