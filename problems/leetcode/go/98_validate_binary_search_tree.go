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

type P98TreeNode struct {
	Val   int
	Left  *P98TreeNode
	Right *P98TreeNode
}

func isValidBST(root *P98TreeNode) bool {
	return dfsP98(root, math.MinInt, math.MaxInt)
}

func dfsP98(node *P98TreeNode, min, max int) bool {
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
		root   *P98TreeNode
		expect bool
	}{
		"case-1": {
			root:   p98BSTCase1(),
			expect: true,
		},
		"case-2": {
			root:   p98BSTCase2(),
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

func p98BSTCase1() *P98TreeNode {
	root := &P98TreeNode{Val: 2}
	root.Left = &P98TreeNode{Val: 1}
	root.Right = &P98TreeNode{Val: 3}
	return root
}

func p98BSTCase2() *P98TreeNode {
	root := &P98TreeNode{Val: 5}
	root.Left = &P98TreeNode{Val: 1}
	root.Right = &P98TreeNode{Val: 4}
	root.Right.Left = &P98TreeNode{Val: 3}
	root.Right.Right = &P98TreeNode{Val: 6}
	return root
}
