package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/*
 * Problem 			: Binary Tree Inorder Traversal
 * Topics           : Stack, Tree, Depth-First Search, Binary Tree
 * Level            : Easy
 * URL              : https://leetcode.com/problems/binary-tree-inorder-traversal
 * Description      :
 * Examples         :
 * 					Example 1:
 * 					Input: root = [1,null,2,3]
 * 					Output: [1,3,2]
 *
 * 					Example 2:
 * 					Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]
 * 					Output: [4,2,6,5,7,1,3,9,8]
 *
 * 					Example 3:
 * 					Input: root = []
 * 					Output: []
 *
 * 					Example 4:
 * 					Input: root = [1]
 * 					Output: [1]
 */

func inorderTraversal(root *TreeNode) []int {
	// Solution using recursive:
	// nodeValues := []int{}
	// return btreeInorderTraverse(root, nodeValues)

	// Solution using iterative:
	return btreeInordertraverseIterative(root)
}

func btreeInorderTraverseRecursive(root *TreeNode, nodeValues []int) []int {
	if root == nil {
		return nodeValues
	}

	nodeValues = btreeInorderTraverseRecursive(root.Left, nodeValues)
	nodeValues = append(nodeValues, root.Val)
	nodeValues = btreeInorderTraverseRecursive(root.Right, nodeValues)

	return nodeValues
}

func btreeInordertraverseIterative(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}
	current := root

	for current != nil || len(stack) > 0 {
		// push all nodes in the left into stack start from current
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		// pop the top node from the stack
		stackLen := len(stack)
		current = stack[stackLen-1]
		stack = stack[:stackLen-1]

		// append current value
		result = append(result, current.Val)

		// go to right node
		current = current.Right
	}

	return result
}

func RunTestBtreeInorderTraversal() {
	testCases := map[string]struct {
		root   *TreeNode
		expect []int
	}{
		"case-1": {
			root:   mockInputP94Case1(),
			expect: []int{1, 3, 2},
		},
		"case-2": {
			root:   mockInputP94Case2(),
			expect: []int{4, 2, 6, 5, 7, 1, 3, 9, 8},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		result := inorderTraversal(testCase.root)
		if !cmp.EqualSlices(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")

	}
	fmt.Printf("\n✅ All tests passed!\n")
}

func mockInputP94Case1() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	return root
}

func mockInputP94Case2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Left.Right.Left = &TreeNode{Val: 6}
	root.Left.Right.Right = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 8}
	root.Right.Right.Left = &TreeNode{Val: 9}
	return root
}
