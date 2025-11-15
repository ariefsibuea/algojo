package main

import (
	"fmt"
	"os"

	"github.com/ariefsibuea/algojo/libs/go/cmp"
)

/*
 * Problem 			: Binary Tree Level Order Traversal
 * Topics           : Tree, Breadth-First Search, Binary Tree
 * Level            : Medium
 * URL              : https://leetcode.com/problems/binary-tree-level-order-traversal
 * Description      : Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e.,
 * 					  from left to right, level by level).
 * Examples         :
 * 					  Example 1:
 * 					  Input: root = [3,9,20,null,null,15,7]
 * 					  Output: [[3],[9,20],[15,7]]
 *
 * 					  Example 2:
 * 					  Input: root = [1]
 * 					  Output: [[1]]
 *
 * 					  Example 3:
 * 					  Input: root = []
 * 					  Output: []
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var (
		result           = [][]int{}
		queue            = []*TreeNode{root}
		currentNode      *TreeNode
		currentLevelSize = 0
	)

	for len(queue) != 0 {
		currentLevelSize = len(queue)
		currentLevelValues := []int{}

		for i := 0; i < currentLevelSize; i++ {
			// dequeue
			currentNode = queue[0]
			queue[0] = nil // release space usage
			queue = queue[1:]

			currentLevelValues = append(currentLevelValues, currentNode.Val)

			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}

		result = append(result, currentLevelValues)
	}

	return result
}

func RunTestBinaryTreeLevelOrderTraversal() {
	testCases := map[string]struct {
		root   *TreeNode
		expect [][]int
	}{
		"case-1": {
			root: mockInputP102Case1(),
			expect: [][]int{
				{3}, {9, 20}, {15, 7},
			},
		},
		"case-2": {
			root: mockInputP102Case2(),
			expect: [][]int{
				{1},
			},
		},
		"case-3": {
			root:   nil,
			expect: [][]int{},
		},
	}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)
		result := levelOrder(testCase.root)
		if !cmp.EqualSlices(result, testCase.expect) {
			fmt.Printf("=== FAILED: expect = %v - got = %v\n", testCase.expect, result)
			os.Exit(1)
		}
		fmt.Printf("=== PASSED\n")
	}

	fmt.Printf("\n✅ All tests passed!\n")
}

func mockInputP102Case1() *TreeNode {
	node7 := &TreeNode{Val: 7}
	node15 := &TreeNode{Val: 15}
	node20 := &TreeNode{Val: 20, Left: node15, Right: node7}
	node9 := &TreeNode{Val: 9}
	return &TreeNode{Val: 3, Left: node9, Right: node20}
}

func mockInputP102Case2() *TreeNode {
	return &TreeNode{Val: 1}
}
