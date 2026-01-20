package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/format"
)

/*
 * Problem 			: Maximum Depth of Binary Tree
 * Topics           : Tree, Depth-First Search, Breadth-First Search, Binary Tree
 * Level            : Easy
 * URL              : https://leetcode.com/problems/maximum-depth-of-binary-tree
 * Description      : Given the root of a binary tree, return its maximum depth. A binary tree's maximum depth is the 
 * 					  number of nodes along the longest path from the root node down to the farthest leaf node.
 * Examples         :
 * 					  Example 1:
 * 					  Input: root = [3,9,20,null,null,15,7]
 * 					  Output: 3
 *
 * 					  Example 2:
 * 					  Input: root = [1,null,2]
 * 					  Output: 2
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	return 1 + max(left, right)
}

func RunTestMaximumDepthOfBinaryTree() {
	// TODO: complete the test function
	testCases := map[string]struct{}{}

	for name, testCase := range testCases {
		fmt.Printf("RUN %s\n", name)

		fmt.Println(testCase)
		// result := twoSum(testCase.nums, testCase.target)
		// format.PrintInput(map[string]interface{}{"input-1": testCase.Input1})

		// if !EqualSlices(result, testCase.expect) {
		//  format.PrintFailed("expect = %v - got = %v\n", testCase.expect, result)
		// 	os.Exit(1)
		// }
		format.PrintSuccess("test case '%s' passed", name)
	}

	fmt.Printf("\n✅ All tests passed!\n")
}
