package main

import (
	"fmt"
	"math"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/**
 * LeetCode Problem : Valide Binary Search Tree
 * Topic            : Tree, Depth-First Search, Binary Search Tree, Binary Tree
 * Level            : Medium
 * URL              : https://leetcode.com/problems/validate-binary-search-tree
 */

func isValidBST(root *TreeNode) bool {
	return dfsP98(root, math.MinInt, math.MaxInt)
}

func dfsP98(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	nodeVal := node.Val
	if nodeVal <= min || nodeVal >= max {
		return false
	}

	return dfsP98(node.Left, min, nodeVal) && dfsP98(node.Right, nodeVal, max)
}

func RunTestIsValidBST() {
	testCases := map[string]struct {
		root   *TreeNode
		expect bool
	}{
		"case-1": {
			root:   mockInputP98Case1(),
			expect: true,
		},
		"case-2": {
			root:   mockInputP98Case2(),
			expect: false,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := isValidBST(testCase.root)
		if !cmp.EqualBooleans(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")

	}
	fmt.Printf("\n✅ All tests passed!\n")
}

func mockInputP98Case1() *TreeNode {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	return root
}

func mockInputP98Case2() *TreeNode {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 6}
	return root
}
