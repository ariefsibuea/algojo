package main

/*
LeetCode Problem : Balanced Binary Tree
Topic            : Tree, Depth-First Search, Binary Tree
Level            : Easy
URL              : https://leetcode.com/problems/balanced-binary-tree/
Description      : Given a binary tree, determine if it is height-balanced. For this problem, a height-balanced binary
        tree is defined as a binary tree in which the left and right subtrees of every node differ in height by no more
        than 1.
Examples         :
        Example 1:
        Input: root = [3,9,20,null,null,15,7]
        Output: true

        Example 2:
        Input: root = [1,2,2,3,3,null,null,4,4]
        Output: false

        Example 3:
        Input: root = []
        Output: true
*/

import (
	"fmt"
	"os"
)

type P110TreeNode struct {
	Val   int
	Left  *P110TreeNode
	Right *P110TreeNode
}

func isBalancedBinaryTree(root *P110TreeNode) bool {
	height := 0
	return isBalancedHeight(root, &height)
}

func isBalancedHeight(root *P110TreeNode, height *int) bool {
	leftHeight := 0
	rightHeight := 0

	if root == nil {
		*height = 0
		return true
	}

	left := isBalancedHeight(root.Left, &leftHeight)
	right := isBalancedHeight(root.Right, &rightHeight)

	*height = max(leftHeight, rightHeight) + 1

	if Abs(leftHeight-rightHeight) <= 1 {
		return left && right
	}

	return false
}

func RunTestIsBalancedBinaryTree() {
	bTreeCase1 := &P110TreeNode{
		Val: 3,
		Left: &P110TreeNode{
			Val: 9,
		},
		Right: &P110TreeNode{
			Val:   20,
			Left:  &P110TreeNode{Val: 15},
			Right: &P110TreeNode{Val: 7},
		},
	}
	bTreeCase2 := &P110TreeNode{
		Val: 1,
		Left: &P110TreeNode{
			Val: 2,
			Left: &P110TreeNode{
				Val:   3,
				Left:  &P110TreeNode{Val: 4},
				Right: &P110TreeNode{Val: 4},
			},
			Right: &P110TreeNode{Val: 3},
		},
		Right: &P110TreeNode{Val: 2},
	}

	testCases := map[string]struct {
		root   *P110TreeNode
		expect bool
	}{
		"case-1": {
			root:   bTreeCase1,
			expect: true,
		},
		"case-2": {
			root:   bTreeCase2,
			expect: false,
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := isBalancedBinaryTree(testCase.root)
		if !EqualBooleans(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
