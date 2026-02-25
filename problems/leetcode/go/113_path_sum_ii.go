package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/format"
)

func init() {
	register("PathSumII", RunTestPathSumII)
}

/*
 * Problem 			: Path Sum II
 * Topics           : Backtracking, Tree, Depth-First Search, Binary Tree
 * Level            : Medium
 * URL              : https://leetcode.com/problems/path-sum-ii
 * Description      : Given the root of a binary tree and an integer targetSum, return all root-to-leaf paths where the
 * 					  sum of the node values in the path equals targetSum. Each path should be returned as a list of
 * 					  the node values, not node references. A leaf is a node with no children.
 *
 * 					  Constraints:
 * 					  - The number of nodes in the tree is in the range [0, 5000].
 * 					  - -1000 <= Node.val <= 1000
 * 					  - -1000 <= targetSum <= 1000
 *
 * Examples         :
 * 					  Example 1:
 * 					  Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
 * 					  Output: [[5,4,11,2],[5,8,4,5]]
 * 					  Explanation: There are two paths whose sum equals targetSum:
 * 					  5 + 4 + 11 + 2 = 22
 * 					  5 + 8 + 4 + 5 = 22
 *
 * 					  Example 2:
 * 					  Input: root = [1,2,3], targetSum = 5
 * 					  Output: []
 *
 * 					  Example 3:
 * 					  Input: root = [1,2], targetSum = 0
 * 					  Output: []
 */

func pathSum(root *TreeNode, targetSum int) [][]int {
	var result = make([][]int, 0)
	var currentPath = make([]int, 0)

	var dfs func(node *TreeNode, targetSum int)

	dfs = func(node *TreeNode, targetSum int) {
		if node == nil {
			return
		}

		currentPath = append(currentPath, node.Val)

		if node.Left == nil && node.Right == nil && node.Val == targetSum {
			temp := make([]int, len(currentPath))
			copy(temp, currentPath)
			result = append(result, temp)
		} else {
			dfs(node.Left, targetSum-node.Val)
			dfs(node.Right, targetSum-node.Val)
		}

		currentPath = currentPath[:len(currentPath)-1]
	}

	dfs(root, targetSum)

	return result
}

func RunTestPathSumII() {
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
