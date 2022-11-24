/* Range Sum of BST
Source		: https://leetcode.com/problems/range-sum-of-bst/
Level		: Easy
Description	: Given the root node of a binary search tree and two integers low and high, return the sum of values of all
			nodes with a value in the inclusive range [low, high].

Example 1:
Input: root = [10,5,15,3,7,null,18], low = 7, high = 15
Output: 32
Explanation: Nodes 7, 10, and 15 are in the range [7, 15]. 7 + 10 + 15 = 32.

Example 2:
Input: root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
Output: 23
Explanation: Nodes 6, 7, and 10 are in the range [6, 10]. 6 + 7 + 10 = 23.
*/

package leetcode

func traverseTree(node *TreeNode, low, hight int) int {
	if node == nil {
		return 0
	}

	leftSum, rightSum := 0, 0
	if node.Val >= low {
		leftSum = traverseTree(node.Left, low, hight)
	}
	if node.Val <= hight {
		rightSum = traverseTree(node.Right, low, hight)
	}

	if node.Val >= low && node.Val <= hight {
		return node.Val + leftSum + rightSum
	}
	return leftSum + rightSum
}

func RangeSumBST(root *TreeNode, low int, high int) int {
	return traverseTree(root, low, high)
}
