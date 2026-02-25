package main

import (
	"fmt"

	"github.com/ariefsibuea/algojo/libs/go/format"
)

func init() {
	register("PathSum", RunTestPathSum)
}

/*
 * Problem 			: Path Sum
 * Topics           : Tree, Depth-First Search, Binary Tree
 * Level            : Easy
 * URL              : https://leetcode.com/problems/path-sum
 * Description      : Given the root of a binary tree and an integer targetSum, return true if the tree has a
 * 					  root-to-leaf path such that adding up all the values along the path equals targetSum.
 * 					  A leaf is a node with no children.
 *
 * 					  Constraints:
 * 					  - The number of nodes in the tree is in the range [0, 5000].
 * 					  - -1000 <= Node.val <= 1000
 * 					  - -1000 <= targetSum <= 1000
 *
 * Examples         :
 * 					  Example 1:
 * 					  Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
 * 					  Output: true
 * 					  Explanation: The root-to-leaf path with the target sum is 5 -> 4 -> 11 -> 2.
 *
 * 					  Example 2:
 * 					  Input: root = [1,2,3], targetSum = 5
 * 					  Output: false
 * 					  Explanation: There are two root-to-leaf paths:
 * 					  (1 --> 2): sum = 3
 * 					  (1 --> 3): sum = 4
 * 					  There is no root-to-leaf path with sum = 5.
 *
 * 					  Example 3:
 * 					  Input: root = [], targetSum = 0
 * 					  Output: false
 * 					  Explanation: Since the tree is empty, there are no root-to-leaf paths.
 */

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	newTargetSum := targetSum - root.Val
	return hasPathSum(root.Left, newTargetSum) || hasPathSum(root.Right, newTargetSum)
}

func RunTestPathSum() {
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
